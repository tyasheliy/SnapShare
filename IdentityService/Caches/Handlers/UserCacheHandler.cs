using IdentityService.Caches.Interfaces;
using IdentityService.Caches.Statuses;
using Microsoft.Extensions.Caching.Distributed;
using System.Text.Json;

namespace IdentityService.Caches.Handlers
{
    /// <summary>
    /// Handles user caching.
    /// </summary>
    public class UserCacheHandler : ICacheHandler<UserCache>
    {
        private IDistributedCache _cache;

        public UserCacheHandler(IDistributedCache cache)
        {
            _cache = cache;
        }

        public async Task<UserCache?> GetAsync(string key)
        {
            string? userData = await _cache.GetStringAsync(key);
            if (userData == null)
            {
                return null;
            }

            return JsonSerializer.Deserialize<UserCache>(userData);
        }

        public async Task<CacheStatus> CacheAsync(string key, UserCache data)
        {
            string? existingData = await _cache.GetStringAsync(key);
            if (existingData != null)
            {
                return CacheStatuses.Exists;
            }

            try
            {
                await _cache.SetStringAsync(key, JsonSerializer.Serialize(data), new DistributedCacheEntryOptions
                {
                    AbsoluteExpirationRelativeToNow = TimeSpan.FromHours(1)
                });
            }
            catch (Exception ex)
            {
                return CacheStatuses.Error(ex);
            }
            return CacheStatuses.Cached;
        }

        public async Task<CacheStatus> UpdateAsync(string key, UserCache updatedData)
        {
            try
            {
                await _cache.RemoveAsync(key);
                await CacheAsync(key, updatedData);
            }
            catch(Exception ex)
            {
                return CacheStatuses.Error(ex);
            }
            return CacheStatuses.Cached;
        }

        public async Task<CacheStatus> DeleteAsync(string key)
        {
            try
            {
                await _cache.RemoveAsync(key);
            }
            catch(Exception ex)
            {
                return CacheStatuses.Error(ex);
            }
            return CacheStatuses.Deleted;
        }
    }
}
