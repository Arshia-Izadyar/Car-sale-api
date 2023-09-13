package service_errors

const (
	OtpExists     = "otp exits 800"
	OtpUsed       = "otp used 801"
	OtpInvalid    = "otp invalid 802"
	ClaimNotFound = "claim not found"

	// user
	EmailExists    = "email already exits"
	UsernameExists = "Username already exits"
	WrongPassword  = "WrongPassword"

	TokenNotPresent = "no token provided"
	TokenExpired    = "token is expired !"
	TokenInvalid    = "provided token is invalid"
	NotRefreshToken = "provided token is not a refresh token"
	InternalError   = "some thing happened"

	PermissionDenied = "Permission Denied"
)
