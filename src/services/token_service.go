package services

import (
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/service_errors"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	Cfg    *config.Config
	Logger logging.Logger
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		Cfg:    cfg,
		Logger: logging.NewLogger(cfg),
	}
}

func (ts *TokenService) GenerateToken(td *dto.TokenDTO) (*dto.TokenDetail, error) {
	tokenDetail := &dto.TokenDetail{}
	tokenDetail.AccessTokenExpireTime = time.Now().Add(time.Minute * ts.Cfg.Jwt.AccessTokenExpireDuration).Unix()
	tokenDetail.RefreshTokenExpireTime = time.Now().Add(time.Minute * ts.Cfg.Jwt.RefreshTokenExpireDuration).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims[constants.UserIdKey] = td.UserId
	accessTokenClaims[constants.FullNameKey] = td.FullName
	accessTokenClaims[constants.UserNameKey] = td.Username
	accessTokenClaims[constants.PhoneKey] = td.Phone
	accessTokenClaims[constants.EmailKey] = td.Email
	accessTokenClaims[constants.RolesKey] = td.Roles
	accessTokenClaims[constants.ExpKey] = tokenDetail.AccessTokenExpireTime

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, err := tk.SignedString([]byte(ts.Cfg.Jwt.Secret))
	if err != nil {
		return nil, err
	}
	tokenDetail.AccessToken = accessToken

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims[constants.UserIdKey] = td.UserId
	refreshTokenClaims[constants.FullNameKey] = td.FullName
	refreshTokenClaims[constants.UserNameKey] = td.Username
	refreshTokenClaims[constants.PhoneKey] = td.Phone
	refreshTokenClaims[constants.EmailKey] = td.Email
	refreshTokenClaims[constants.RolesKey] = td.Roles
	refreshTokenClaims[constants.ExpKey] = tokenDetail.AccessTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, err := rt.SignedString([]byte(ts.Cfg.Jwt.RefreshSecret))
	if err != nil {
		return nil, err
	}
	tokenDetail.RefreshToken = refreshToken
	return tokenDetail, nil
}

func (ts *TokenService) ValidateToken(token string) (*jwt.Token, error) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
		}
		return []byte(ts.Cfg.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (ts *TokenService) GetClaims(token string) (map[string]interface{}, error) {
	claimMap := make(map[string]interface{})
	verifiedToken, err := ts.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifiedToken.Claims.(jwt.MapClaims)
	if ok && verifiedToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimNotFound}
}
