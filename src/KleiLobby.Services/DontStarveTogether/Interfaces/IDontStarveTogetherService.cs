using KleiLobby.Domain.DontStarveTogether;

namespace KleiLobby.Services.DontStarveTogether.Interfaces
{
    public interface IDontStarveTogetherService
    {
        public Task<IEnumerable<ServerInfo>> GetAllAsync();
        public Task<IEnumerable<ServerInfo>> GetByHostAndNameAsync(string host, string name);
        public Task<ServerInfo?> GetByRowIdAsync(string rowId);
    }
}
