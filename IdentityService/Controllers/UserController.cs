using IdentityService.Caches;
using IdentityService.Caches.Handlers;
using IdentityService.Caches.Statuses;
using IdentityService.Data;
using IdentityService.DTO;
using IdentityService.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Caching.Distributed;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text.Json;

namespace IdentityService.Controllers
{
    /// <summary>
    /// User controller manages users of the application.
    /// </summary>
    [ApiController]
    [Route("users")]
    public class UserController : ControllerBase
    {
        private UserManager<User> _userManager;
        private UserCacheHandler _userCaches;

        public UserController(UserManager<User> userManager, UserCacheHandler userCaches)
        {
            _userManager = userManager;
            _userCaches = userCaches;
        }

        /// <summary>
        /// Reads user authorization token to retrieve user identifier.
        /// </summary>
        /// <returns>
        /// Status code 200 with user identifier is succeeded.
        /// Status code 400 if token is invalid.
        /// Status code 500 if token's claims has no user identifier.
        /// </returns>
        [HttpGet]
        [Authorize]
        public async Task<IActionResult> Get(string? intention)
        {
            string token = Request.Headers.Authorization.ToString().Remove(0, 7);

            var handler = new JwtSecurityTokenHandler();
            if (!handler.CanReadToken(token))
            {
                return BadRequest(new { Message = "Invalid token" });
            }

            var jwt = handler.ReadJwtToken(token);
            if (!jwt.Claims.Any())
            {
                return BadRequest(new { Message = "Invalid token" });
            }

            string id = jwt.Claims.FirstOrDefault(c => c.Type == "userId").Value;
            if (id.Equals(string.Empty) || id.Equals(null))
            {
                return StatusCode(500, new { Message = "Invalid user identifier" });
            }

            switch (intention)
            {
                case "identifier":
                    return Ok(new { Id = id });

                case "data":
                    UserCache? cache = await _userCaches.GetAsync(id);
                    if (cache != null)
                    {
                        return Ok(cache);
                    }

                    User? user = await _userManager.FindByIdAsync(id);
                    if (user == null)
                    {
                        return BadRequest(new { Message = "User is not found" });
                    }

                    return Ok(new UserCache(user));

                default:
                    return Ok(new { Id = id });
            }
        }

        /// <summary>
        /// This method tries to registrate a new user with the given data.
        /// </summary>
        /// <param name="userData">User registration data.</param>
        /// <returns>
        /// Status code 204 if registration succeeded.
        /// Status code 400 with errors if not.
        /// </returns>
        [HttpPost]
        public async Task<IActionResult> Post([FromBody] UserRegistrationDTO userData)
        {
            User user = new User
            {
                Email = userData.Email,
                UserName = userData.UserName,
            };

            var status = await _userManager.CreateAsync(user, userData.Password);
            if (!status.Succeeded)
            {
               return BadRequest(new { Message = "Failed to create an user" });
            }

            // TODO: Add logger.
            var userCache = new UserCache(user);
            var cacheStatus = await _userCaches.CacheAsync(user.Id.ToString(), userCache);

            return NoContent();
        }

        /// <summary>
        /// Changes user's password.
        /// </summary>
        /// <param name="userData">User data required for password changing.</param>
        /// <returns>
        /// Status code 204 if succeeded.
        /// Status code 400 if there were any errors.
        /// </returns>
        [HttpPatch]
        public async Task<IActionResult> Patch([FromBody] UserChangePasswordDTO userData)
        {
            User user = await _userManager.FindByNameAsync(userData.UserName);

            if (user == null)
            {
                return BadRequest(new { Message = "User is not found" });
            }

            var status = await _userManager.ChangePasswordAsync(user, userData.OldPassword, userData.Password);
            if (!status.Succeeded)
            {
                return BadRequest(new
                {
                    Message = "Failed to change password",
                    Errors = status.Errors.ToList()
                });
            }

            return NoContent();
        }
    }
}
