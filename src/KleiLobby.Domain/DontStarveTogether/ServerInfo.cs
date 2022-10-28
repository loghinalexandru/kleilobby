using Newtonsoft.Json;

namespace KleiLobby.Domain.DontStarveTogether
{
    public sealed class ServerInfo
    {
        [JsonProperty("__addr")]
        public string? Address { get; set; }
        [JsonProperty("__lastping")]
        public long LastPing { get; set; }
        [JsonProperty("__rowId")]
        public string? RowId { get; set; }
        [JsonProperty("host")]
        public string? HostKU { get; set; }
        [JsonProperty("name")]
        public string? Name { get; set; }
        [JsonProperty("password")]
        public bool? Password { get; set; }
        [JsonProperty("connected")]
        public short? Connected { get; set; }
        [JsonProperty("season")]
        public string? Season { get; set; }
        [JsonProperty("serverpaused")]
        public bool? ServerPaused { get; set; }
    }
}
