using IdentityService.Caches.Handlers;
using IdentityService.Data;
using IdentityService.Models;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;
using StackExchange.Redis;
using System.Net.Sockets;
using System.Text;

internal class Program
{
    private static int Main(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);

        try
        {
            using var dbClient = new TcpClient("host.docker.internal", 5001);
            using var cacheClient = new TcpClient("host.docker.internal", 5002);
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex);
            return 1;
        }

        var services = builder.Services;

        //adding services
        services.AddCors();
        services.AddControllers();
        services.AddStackExchangeRedisCache(o =>
        {
            o.Configuration = $"host.docker.internal:5002, password={Environment.GetEnvironmentVariable("IDENTITY_REDIS_PASSWORD")}";
            o.InstanceName = "identity.";
        });
        services.AddDbContext<DataContext>(o => o.UseNpgsql($"Server=host.docker.internal; Port=5001; Password={Environment.GetEnvironmentVariable("IDENTITY_PG_PASSWORD")}; Database={Environment.GetEnvironmentVariable("IDENTITY_PG_DB")}; Userid={Environment.GetEnvironmentVariable("IDENTITY_PG_USER")}"));
        services.AddIdentityCore<User>(o =>
        {
            //development
            o.Password = new PasswordOptions
            {
                RequireDigit = false,
                RequiredLength = 1,
                RequiredUniqueChars = 0,
                RequireLowercase = false,
                RequireNonAlphanumeric = false,
                RequireUppercase = false
            };
        })
            .AddEntityFrameworkStores<DataContext>()
            .AddSignInManager();

        services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
            .AddJwtBearer(o =>
            {
                //development
                o.RequireHttpsMetadata = false;

                o.TokenValidationParameters = new TokenValidationParameters()
                {
                    ValidateAudience = true,
                    ValidateIssuer = true,
                    ValidateIssuerSigningKey = true,
                    ValidateLifetime = true,

                    ValidAudience = builder.Configuration.GetValue<string>("TokenAudience"),
                    ValidIssuer = builder.Configuration.GetValue<string>("TokenIssuer"),
                    IssuerSigningKey = new SymmetricSecurityKey(Encoding.ASCII.GetBytes(Environment.GetEnvironmentVariable("TOKEN_SECRET_KEY")))
                };
            });
        services.AddAuthorization();

        services.AddScoped<UserCacheHandler>();

        var app = builder.Build();

        //using services
        app.UseAuthentication();
        app.UseRouting();
        app.UseCors(b => b.AllowAnyOrigin().AllowAnyMethod().AllowAnyHeader());
        app.UseAuthorization();
        app.UseEndpoints(endpoints => endpoints.MapControllers());
        

        app.Run();
        return 0;
    }
}