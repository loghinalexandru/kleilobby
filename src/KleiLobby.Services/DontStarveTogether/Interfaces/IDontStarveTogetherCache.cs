using KleiLobby.Domain.DontStarveTogether;
using KleiLobby.Services.DontStarveTogether.Constants;

namespace KleiLobby.Services.DontStarveTogether.Interfaces
{
    public interface IDontStarveTogetherCache
    {
        public string GetServerRowId(LobbyListEnum regionKey, string host, string serverName);

        public void RemoveKey(string key);

        public Task<RequestWrapper> GetRequestWrapper(LobbyListEnum regionKey);

        public Task<bool> SetRequestWrapper(LobbyListEnum regionKey, RequestWrapper request);
    }
}
