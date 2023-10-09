namespace IdentityService.DTO
{
    /// <summary>
    /// Represents a pair of username and password of the user.
    /// </summary>
    public struct UserCredentialsDTO
    {
        public string UserName { get; set; }
        public string Password { get; set; }
    }
}
