using KleiLobby.Domain.DontStarveTogether;
using KleiLobby.Services.DontStarveTogether.Constants;
using KleiLobby.Services.DontStarveTogether.Interfaces;
using Microsoft.Extensions.Caching.Memory;

namespace KleiLobby.Services.DontStarveTogether
{
    public sealed class DontStarveTogetherCache : IDontStarveTogetherCache
    {
        private readonly IMemoryCache _cache;
        private readonly TimeSpan _cacheExpirationTime = TimeSpan.FromMinutes(5);
        private readonly MemoryCacheEntryOptions _entryOptions = new MemoryCacheEntryOptions { Priority = CacheItemPriority.NeverRemove };

        public DontStarveTogetherCache(IMemoryCache cache)
        {
            _cache = cache;
        }

        public Task<RequestWrapper> GetRequestWrapper(LobbyListEnum regionKey)
        {
            _cache.TryGetValue<RequestWrapper>(Enum.GetName(regionKey), out var lobbyList);

            return Task.FromResult(lobbyList);
        }

        public string GetServerRowId(LobbyListEnum regionKey, string host, string serverName)
        {
            _cache.TryGetValue<string>(GetServerRowKey(regionKey, host, serverName), out var lobbyList);

            return lobbyList;
        }

        public Task<bool> SetRequestWrapper(LobbyListEnum regionKey, RequestWrapper request)
        {
            try
            {
                _cache.Set(Enum.GetName(regionKey), request, _cacheExpirationTime);
                SetServersRowID(regionKey, request);

                return Task.FromResult(true);
            }
            catch (Exception)
            {
                return Task.FromResult(false);
            }
        }

        private void SetServersRowID(LobbyListEnum regionKey, RequestWrapper request)
        {
            foreach (var entry in request.Lobby?.Where(x => !string.IsNullOrWhiteSpace(x.HostKU) && !string.IsNullOrWhiteSpace(x.Name)) ?? Enumerable.Empty<ServerInfo>())
            {
                _cache.Set(GetServerRowKey(regionKey, entry.HostKU!, entry.Name!), entry.RowId, _entryOptions);
            }
        }

        private string GetServerRowKey(LobbyListEnum regionKey, string host, string serverName)
        {
            return string.Join(":", Enum.GetName(regionKey), host, serverName);
        }

        public void RemoveKey(string key)
        {
            _cache.Remove(key);
        }
    }
}
