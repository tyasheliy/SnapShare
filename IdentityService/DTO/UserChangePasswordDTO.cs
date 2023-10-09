namespace IdentityService.DTO
{
    /// <summary>
    /// Represents required data for changing user password.
    /// </summary>
    public struct UserChangePasswordDTO
    {
        public string UserName { get; set; }
        public string OldPassword { get; set; }
        public string Password { get; set; }
    }
}
