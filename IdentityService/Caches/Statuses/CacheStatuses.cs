namespace IdentityService.Caches.Statuses
{
    public static class CacheStatuses
    {
        public static readonly CacheStatus Cached = new CacheStatus("cached");
        public static readonly CacheStatus Exists = new CacheStatus("exists");
        public static readonly CacheStatus Deleted = new CacheStatus("deleted");
        
        public static CacheStatus Error(Exception ex)
        {
            return new CacheStatus(ex.ToString());
        }

        public static bool IsCached(CacheStatus status)
        {
            return status == Cached;
        }
    }
}
