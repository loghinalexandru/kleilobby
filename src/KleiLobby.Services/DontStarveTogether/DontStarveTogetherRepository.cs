using KleiLobby.Domain.DontStarveTogether;
using KleiLobby.Services.DontStarveTogether.Interfaces;
using Newtonsoft.Json;
using System.Net.Http.Headers;

namespace KleiLobby.Services.DontStarveTogether
{
    public sealed class DontStarveTogetherRepository : IDontStarveTogetherRepository
    {
        private readonly IHttpClientFactory _httpClientFactory;
        private readonly IContextResolver _contextResolver;

        public DontStarveTogetherRepository(
            IHttpClientFactory httpClientFactory,
            IContextResolver contextResolver)
        {
            _httpClientFactory = httpClientFactory;
            _contextResolver = contextResolver;
        }

        public async Task<RequestWrapper?> GetAll()
        {
            var client = _httpClientFactory.CreateClient(nameof(DontStarveTogetherService));
            var region = _contextResolver.GetRawRegion();
            var token = _contextResolver.GetToken();

            var request = new HttpRequestMessage(HttpMethod.Post, $@"https://lobby-{region}.klei.com/lobby/read");
            request.Content = new StringContent("{\"__gameId\": \"DST\",\"__token\": \"token_to_replace\"}".Replace("token_to_replace", token));
            request.Content.Headers.ContentType = new MediaTypeHeaderValue("application/x-www-form-urlencoded");

            var response = await client.SendAsync(request);

            if (!response.IsSuccessStatusCode)
            {
                return null;
            }

            return JsonConvert.DeserializeObject<RequestWrapper>(await response.Content.ReadAsStringAsync());
        }

        public async Task<ServerInfo?> GetByRowId(string rowId)
        {
            var client = _httpClientFactory.CreateClient(nameof(DontStarveTogetherService));
            var region = _contextResolver.GetRawRegion();
            var token = _contextResolver.GetToken();

            var request = new HttpRequestMessage(HttpMethod.Post, $@"https://lobby-{region}.klei.com/lobby/read");
            request.Content = new
                StringContent("{\"__gameId\": \"DST\",\"__token\": \"token_to_replace\", \"query\":{\"__rowId\":\"rowId_to_replace\"}}}"
                .Replace("token_to_replace", token)
                .Replace("rowId_to_replace", rowId));

            request.Content.Headers.ContentType = new MediaTypeHeaderValue("application/x-www-form-urlencoded");

            var response = await client.SendAsync(request);

            var server = JsonConvert.DeserializeObject<RequestWrapper>(await response.Content.ReadAsStringAsync());

            return server?.Lobby?.FirstOrDefault();
        }
    }
}
