using KleiLobby.Domain.DontStarveTogether;
using KleiLobby.Services.DontStarveTogether.Interfaces;

namespace KleiLobby.Services.DontStarveTogether
{
    public sealed class DontStarveTogetherService : IDontStarveTogetherService
    {
        private readonly IDontStarveTogetherCache _cache;
        private readonly IDontStarveTogetherRepository _repository;
        private readonly IContextResolver _contextResolver;

        public DontStarveTogetherService(
            IDontStarveTogetherCache cache,
            IContextResolver contextResolver,
            IDontStarveTogetherRepository repository)
        {
            _cache = cache;
            _contextResolver = contextResolver;
            _repository = repository;
        }

        public async Task<IEnumerable<ServerInfo>> GetAllAsync()
        {
            var regionKey = _contextResolver.GetLobbyRegion();

            var result = _cache.GetRequestWrapper(regionKey);

            result ??= await _repository.GetAll();

            if (result != null && (result.Lobby?.Any() ?? false))
            {
                _cache.SetRequestWrapper(regionKey, result);
            }

            return result?.Lobby ?? Enumerable.Empty<ServerInfo>();
        }

        public async Task<ServerInfo?> GetByHostAndNameAsync(string host, string name)
        {
            var regionKey = _contextResolver.GetLobbyRegion();
            var serverRowId = _cache.GetServerRowId(regionKey, host, name);

            if (!string.IsNullOrWhiteSpace(serverRowId))
            {
                var serverByRowId = await _repository.GetByRowId(serverRowId);

                if (serverByRowId != null)
                {
                    return serverByRowId;
                }
            }

            var result = await _repository.GetAll();

            if (result != null && (result.Lobby?.Any() ?? false))
            {
                _cache.SetRequestWrapper(regionKey, result);
            }

            return
                result?.Lobby?
                .Where(x => x.HostKU!.Equals(host, StringComparison.InvariantCultureIgnoreCase) && x.Name!.Equals(name, StringComparison.InvariantCultureIgnoreCase))
                .FirstOrDefault();
        }

        public async Task<ServerInfo?> GetByRowIdAsync(string rowId)
        {
            var result = await _repository.GetByRowId(rowId);

            return result;
        }
    }
}
