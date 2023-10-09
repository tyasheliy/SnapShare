namespace IdentityService.Caches.Statuses
{
    public class CacheStatus
    {
        public string Status { get; }

        public CacheStatus(string status) => Status = status;
    }
}
