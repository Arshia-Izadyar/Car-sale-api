package constants

const (
	AdminRoleName           string = "admin"
	DefaultRoleName         string = "default"
	DefaultUserName         string = "admin"
	DefaultRedisKey                = "otp"
	AuthenTicationHeaderKey        = "Authorization"

	// claims
	UserIdKey   = "UserId"
	FullNameKey = "FullName"
	UserNameKey = "UserName"
	PhoneKey    = "Phone"
	EmailKey    = "Email"
	RolesKey    = "Roles"
	ExpKey      = "Exp"
	AccessType  = "access_token"
	RefreshType = "refresh_token"
)
