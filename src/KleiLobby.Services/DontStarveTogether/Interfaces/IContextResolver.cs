using KleiLobby.Services.DontStarveTogether.Constants;

namespace KleiLobby.Services.DontStarveTogether.Interfaces
{
    public interface IContextResolver
    {
        public LobbyListEnum GetLobbyRegion();

        public string GetRawRegion();

        public string GetToken();
    }
}
