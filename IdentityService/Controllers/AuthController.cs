using IdentityService.Caches;
using IdentityService.Caches.Handlers;
using IdentityService.DTO;
using IdentityService.Models;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Caching.Distributed;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using System.Text.Json;

namespace IdentityService.Controllers
{
    /// <summary>
    /// Auth controller handles users' authentication process.
    /// </summary>
    [ApiController]
    [Route("auth")]
    public class AuthController : ControllerBase
    {
        private SignInManager<User> _signInManager;
        private IConfiguration _configuration;
        private UserCacheHandler _userCaches;

        public AuthController(SignInManager<User> signInManager, 
            IConfiguration configuration,
            UserCacheHandler userCashes)
        {
            _signInManager = signInManager;
            _configuration = configuration;
            _userCaches = userCashes;
        }

        /// <summary>
        /// This method authenticates the user with given credentials.
        /// </summary>
        /// <param name="userCredentials">User credentials</param>
        /// <returns>
        /// Status code 200 with token if succeeded.
        /// Status code 400 if there were any errors.
        /// </returns>
        [HttpPost]
        public async Task<IActionResult> Authenticate([FromBody] UserCredentialsDTO userCredentials)
        {
            User user = await _signInManager.UserManager.FindByNameAsync(userCredentials.UserName);
            if (user == null)
            {
                return BadRequest(new { Message = "User with given username is not found" });
            }

            var status = await _signInManager.CheckPasswordSignInAsync(user, userCredentials.Password, false);
            if (!status.Succeeded)
            {
                return BadRequest(new { Message = "Invalid user credentials" });
            }

            List<Claim> claims = new List<Claim>
            {
                new Claim(ClaimTypes.NameIdentifier, user.Id.ToString())
            };
            var securityKey = new SymmetricSecurityKey(Encoding.ASCII.GetBytes(_configuration.GetValue<string>("TokenSecretKey")));

            var token = new JwtSecurityToken(
                issuer: _configuration.GetValue<string>("TokenIssuer"),
                audience: _configuration.GetValue<string>("TokenAudience"),
                signingCredentials: new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256Signature),

                expires: DateTime.Now.AddMinutes(_configuration.GetValue<double>("TokenLifetime")),
                claims: claims
                );

            var userCache = new UserCache(user);
            await _userCaches.CacheAsync(user.Id.ToString(), userCache);

            return Ok(new { Token = new JwtSecurityTokenHandler().WriteToken(token) });
        }
    }
}
