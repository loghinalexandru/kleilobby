using KleiLobby.Services.DontStarveTogether.Http;
using System.Net;

namespace KleiLobby.Extensions
{
    public static class HttpClientsExtension
    {
        public static IServiceCollection AddNamedHttpClients(this IServiceCollection services)
        {
            services.AddHttpClient(HttpClients.GZip)
                .ConfigurePrimaryHttpMessageHandler(() => new HttpClientHandler
                {
                    AutomaticDecompression = DecompressionMethods.GZip | DecompressionMethods.Deflate
                });
            services.AddHttpClient(HttpClients.Default);

            return services;
        }
    }
}
