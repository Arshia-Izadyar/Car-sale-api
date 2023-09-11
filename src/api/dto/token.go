package dto

type TokenDTO struct {
	UserId   int
	FullName string
	Username string
	Phone    string
	Email    string
	Roles    []string
}

type TokenDetail struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpireTime  int64  `json:"access_token_expire"`
	RefreshTokenExpireTime int64  `json:"refresh_token_expire"`
}
