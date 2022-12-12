using KleiLobby.Middleware;
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
            services.AddSwaggerGen();
            services.AddSingleton<GlobalExceptionHandlingMiddleware>();

            services.AddCors(def => def.AddDefaultPolicy(p =>
             {
                 p.AllowAnyHeader();
                 p.AllowAnyMethod();
                 p.AllowAnyOrigin();
             }));

            services.AddScoped<IDontStarveTogetherService, DontStarveTogetherService>();
            services.AddScoped<IDontStarveTogetherRepository, DontStarveTogetherRepository>();
            services.AddSingleton<IContextResolver, ContextResolver>();
            services.AddSingleton<IDontStarveTogetherCache, DontStarveTogetherCache>();
        }

        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseRouting();
            app.UseMiddleware<GlobalExceptionHandlingMiddleware>();
            app.UseCors();

            app.UseSwagger();
            app.UseSwaggerUI();

            app.UseEndpoints(x =>
            {
                x.MapControllers();
            });
        }
    }
}