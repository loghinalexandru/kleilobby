namespace KleiLobby
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var hostBuilder = Host.CreateDefaultBuilder(args);
            hostBuilder.ConfigureWebHostDefaults(webHostBuilder =>
            {
                webHostBuilder.UseStartup<Startup>();
            });

           hostBuilder.Build().Run();
        }
    }
}
