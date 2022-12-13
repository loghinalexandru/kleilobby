using KleiLobby.Domain.DontStarveTogether;
using KleiLobby.Services.DontStarveTogether.Constants;

namespace KleiLobby.Services.DontStarveTogether.Interfaces
{
    public interface IDontStarveTogetherCache
    {
        public string GetServerRowId(LobbyListEnum regionKey, string host, string serverName);

        public RequestWrapper GetRequestWrapper(LobbyListEnum regionKey);

        public bool SetRequestWrapper(LobbyListEnum regionKey, RequestWrapper request);
    }
}
