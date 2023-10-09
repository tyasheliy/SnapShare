using IdentityService.Models;

namespace IdentityService.Caches
{
    /// <summary>
    /// Represents user stored in cache.
    /// </summary>
    public struct UserCache
    {
        public UserCache(User user)
        {
            UserName = user.UserName;
            Email = user.Email;
            IsEmailConfirmed = user.EmailConfirmed;
        }

        public string UserName { get; set; }
        public string Email { get; set; }
        public bool IsEmailConfirmed { get; set; }
    }
}
