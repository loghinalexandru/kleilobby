using KleiLobby.Services.DontStarveTogether;
using KleiLobby.Services.DontStarveTogether.Interfaces;

namespace KleiLobby
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        public void ConfigureServices(IServiceCollection services)
        {
            services.AddControllers();
            services.AddMemoryCache();
            services.AddHttpContextAccessor();
            services.AddHttpClient();
            services.AddMvc();
            services.AddCors(def => def.AddDefaultPolicy(p =>
             {
                 p.AllowAnyHeader();
                 p.AllowAnyMethod();
                 p.AllowAnyOrigin();
             }));
            services.AddScoped<IDontStarveTogetherService, DontStarveTogetherService>();
        }

        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseRouting();
            app.UseCors();
            app.UseEndpoints(x =>
            {
                x.MapControllers();
            });
        }
    }
}