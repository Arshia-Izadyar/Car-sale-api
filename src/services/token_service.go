package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/cache"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
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
	accessTokenClaims[constants.AccessType] = true

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, err := tk.SignedString([]byte(ts.Cfg.Jwt.Secret))
	if err != nil {
		return nil, err
	}
	tokenDetail.AccessToken = accessToken

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims[constants.UserIdKey] = td.UserId
	refreshTokenClaims[constants.ExpKey] = tokenDetail.RefreshTokenExpireTime
	refreshTokenClaims[constants.RefreshType] = true

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
	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		expirationTime := time.Unix(int64(claims[constants.ExpKey].(float64)), 0)
		currentTime := time.Now()

		if currentTime.After(expirationTime) {
			// Token has expired
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
		}
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

func (ts *TokenService) ValidateRefreshToken(rToken string) (*dto.TokenDetail, error) {

	// Parse the refresh token
	token, err := jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		// Provide the key used to sign the token here
		return []byte(ts.Cfg.Jwt.Secret), nil
	})

	// Check if the token is valid and not expired
	if err != nil {
		return nil, errors.New("the provided token is wrong")
	}

	isBlackList := isInBlackList(rToken)
	if err != nil {
		return nil, errors.New("internal error")
	}
	claimMap := token.Claims.(jwt.MapClaims)
	if _, ok := claimMap[constants.RefreshType]; !ok {
		return nil, errors.New("provided token is not a refresh token")
	}
	if !token.Valid || isBlackList {
		return nil, fmt.Errorf("refresh token is not valid")
	}

	tokenExp := claimMap[constants.ExpKey]
	expTime := float64(time.Now().Unix()) - tokenExp.(float64)

	addToBlackList(rToken, expTime)

	db := db.GetDB()
	user := models.User{}
	err = db.Model(&models.User{}).Where("id = ?", claimMap[constants.UserIdKey]).First(&user).Error
	if err != nil {
		return nil, err
	}

	tDTO := &dto.TokenDTO{
		UserId:   user.ID,
		FullName: fmt.Sprintf("%s %s", user.FirstName, user.LastName.String),
		Username: user.Username,
		Phone:    user.PhoneNumber.String,
		Email:    user.Email.String,
	}
	if len(user.UserRoles) > 0 {
		for _, r := range user.UserRoles {
			tDTO.Roles = append(tDTO.Roles, r.Role.Name)
		}
	}
	tk, err := ts.GenerateToken(tDTO)
	if err != nil {
		return nil, err
	}

	return tk, nil
}

func addToBlackList(token string, expTime float64) error {
	rds := cache.GetRedis()
	err := cache.Set(token, true, time.Duration(expTime), rds)
	if err != nil {
		return err
	}
	return nil
}

func isInBlackList(token string) bool {
	rds := cache.GetRedis()
	_, err := cache.Get[bool](token, rds)
	return err == nil
}
