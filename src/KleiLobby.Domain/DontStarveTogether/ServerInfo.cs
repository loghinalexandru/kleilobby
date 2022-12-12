using Newtonsoft.Json;

namespace KleiLobby.Domain.DontStarveTogether
{
    public sealed class ServerInfo
    {
        private readonly string? _rawServerData;

        [JsonProperty("__addr")]
        public string? Address { get; init; }
        [JsonProperty("__lastping")]
        public long LastPing { get; init; }
        [JsonProperty("__rowId")]
        public string? RowId { get; init; }
        [JsonProperty("host")]
        public string? HostKU { get; init; }
        [JsonProperty("name")]
        public string? Name { get; init; }
        [JsonProperty("password")]
        public bool? Password { get; init; }
        [JsonProperty("mods")]
        public bool? Mods { get; init; }
        [JsonProperty("connected")]
        public short? Connected { get; init; }
        [JsonProperty("season")]
        public string? Season { get; init; }
        [JsonProperty("serverpaused")]
        public bool? ServerPaused { get; init; }
        [JsonProperty("data")]
        public string? RawServerData
        {
            get { return _rawServerData; }
            init
            {
                _rawServerData = value;
                ServerData = new ServerData(_rawServerData ?? string.Empty);
            }
        }

        public ServerData? ServerData { get; private init; }
    }
}
