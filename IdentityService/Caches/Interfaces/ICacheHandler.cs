using IdentityService.Caches.Statuses;

namespace IdentityService.Caches.Interfaces
{
    // TODO: Document methods.

    /// <summary>
    /// Represents simple CRUD operations with cache.
    /// </summary>
    public interface ICacheHandler<T> where T : struct
    {
        Task<T?> GetAsync(string key);
        Task<CacheStatus> CacheAsync(string key, T data);
        Task<CacheStatus> UpdateAsync(string key, T updatedData);
        Task<CacheStatus> DeleteAsync(string key);
    }
}
