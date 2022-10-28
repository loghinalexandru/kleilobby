using KleiLobby.Domain.DontStarveTogether;
using KleiLobby.Services.DontStarveTogether.Constants;
using KleiLobby.Services.DontStarveTogether.Interfaces;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Caching.Memory;
using Newtonsoft.Json;
using System.Net.Http.Headers;

namespace KleiLobby.Services.DontStarveTogether
{
    public sealed class DontStarveTogetherService : IDontStarveTogetherService
    {
        private readonly IMemoryCache _cache;
        private readonly IHttpContextAccessor _httpContextAccessor;
        private readonly IHttpClientFactory _httpClientFactory;

        public DontStarveTogetherService(
            IMemoryCache cache,
            IHttpContextAccessor httpContextAccessor,
            IHttpClientFactory httpClientFactory)
        {
            _cache = cache;
            _httpContextAccessor = httpContextAccessor;
            _httpClientFactory = httpClientFactory;
        }

        public async Task<IEnumerable<ServerInfo>> GetAllAsync()
        {
            _cache.TryGetValue<RequestWrapper>(LobbyListKeys.ProdLobby, out var lobbyList);

            if(lobbyList == null)
            {
                lobbyList = await GetLobbyList();
            }

            return lobbyList?.Lobby ?? Enumerable.Empty<ServerInfo>();
        }

        public async Task<IEnumerable<ServerInfo>> GetByHostAndNameAsync(string host, string name)
        {
            _cache.TryGetValue<RequestWrapper>(LobbyListKeys.ProdLobby, out var lobbyList);

            if (lobbyList == null)
            {
                lobbyList = await GetLobbyList();
            }

            return 
                lobbyList?.Lobby?
                .Where(x => x.HostKU!.Equals(host, StringComparison.InvariantCultureIgnoreCase) && x.Name!.Equals(name, StringComparison.InvariantCultureIgnoreCase))
                ?? Enumerable.Empty<ServerInfo>();
        }

        public async Task<ServerInfo?> GetByRowIdAsync(string rowId)
        {
            _cache.TryGetValue<RequestWrapper>(LobbyListKeys.ProdLobby, out var lobbyList);

            if (lobbyList == null)
            {
                lobbyList = await GetLobbyList();
            }

            return
                lobbyList?.Lobby?
                .Where(x => x.RowId!.Equals(rowId, StringComparison.InvariantCultureIgnoreCase))
                .FirstOrDefault();
        }

        private async Task<RequestWrapper?> GetLobbyList()
        {
            var client = _httpClientFactory.CreateClient(nameof(DontStarveTogetherService));
            var region = _httpContextAccessor.HttpContext.Request.Query["region"].FirstOrDefault();
            var token = _httpContextAccessor.HttpContext.Request.Query["token"].FirstOrDefault();

            var request = new HttpRequestMessage(HttpMethod.Post, $@"https://lobby-{region}.klei.com/lobby/read");
            request.Content = new StringContent("{\"__gameId\": \"DST\",\"__token\": \"token_to_replace\"}".Replace("token_to_replace", token));
            request.Content.Headers.ContentType = new MediaTypeHeaderValue("application/x-www-form-urlencoded");

            var response = await client.SendAsync(request);

            if (response.IsSuccessStatusCode)
            {
                return _cache.Set(LobbyListKeys.ProdLobby, JsonConvert.DeserializeObject<RequestWrapper>(await response.Content.ReadAsStringAsync()), TimeSpan.FromMinutes(5));
            }

            return null;
        }
    }
}
