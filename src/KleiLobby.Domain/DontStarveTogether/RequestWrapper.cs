using Newtonsoft.Json;

namespace KleiLobby.Domain.DontStarveTogether
{
    public sealed class RequestWrapper
    {
        [JsonProperty("GET")]
        public List<ServerInfo>? Lobby { get; set; }
    }
}
