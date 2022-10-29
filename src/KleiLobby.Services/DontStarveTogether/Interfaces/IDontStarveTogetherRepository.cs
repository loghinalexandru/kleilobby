using KleiLobby.Domain.DontStarveTogether;

namespace KleiLobby.Services.DontStarveTogether.Interfaces
{
    public interface IDontStarveTogetherRepository
    {
        public Task<RequestWrapper?> GetAll();
        public Task<ServerInfo?> GetByRowId(string rowId);
    }
}
