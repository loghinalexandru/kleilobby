﻿using KleiLobby.Services.DontStarveTogether.Constants;
using KleiLobby.Services.DontStarveTogether.Interfaces;
using Microsoft.AspNetCore.Http;

namespace KleiLobby.Services.DontStarveTogether
{
    public sealed class ContextResolver : IContextResolver
    {
        private readonly IHttpContextAccessor _httpContextAccessor;

        public ContextResolver(IHttpContextAccessor httpContextAccessor)
        {
            _httpContextAccessor = httpContextAccessor;
        }

        public LobbyListEnum GetLobbyRegion()
        {
            var region = _httpContextAccessor.HttpContext.Request.Query["region"].FirstOrDefault();

            switch (region)
            {
                case "ap-east-1":
                    return LobbyListEnum.AsianPacificEastLobby;
                case "ap-southeast-1":
                    return LobbyListEnum.AsianPacificSouthEastLobby;
                case "us-east-1":
                    return LobbyListEnum.UsEastLobby;
                case "eu-central-1":
                    return LobbyListEnum.EuCentralLobby;
                default:
                    return LobbyListEnum.Unknown;
            }
        }

        public string GetRawRegion()
        {
            return _httpContextAccessor.HttpContext.Request.Query["region"].FirstOrDefault() ?? string.Empty;
        }

        public string GetToken()
        {
            return _httpContextAccessor.HttpContext.Request.Query["token"].FirstOrDefault() ?? string.Empty;
        }
    }
}
