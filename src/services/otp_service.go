package services

import (
	"fmt"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/cache"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/service_errors"
	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	Logger logging.Logger
	Cfg    *config.Config
	Redis  *redis.Client
}

func NewOTPService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{
		Logger: logger,
		Redis:  redis,
		Cfg:    cfg,
	}
}

func (os *OtpService) SetOtp(mobileNumber, otp string) *service_errors.ServiceError {
	key := fmt.Sprintf("%s:%s", constants.DefaultRedisKey, mobileNumber)
	val := dto.OtpDTO{
		Value: key,
		Used:  false,
	}
	result, err := cache.Get[dto.OtpDTO](key, os.Redis)
	if err == nil && result.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if err == nil && !result.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	}
	err = cache.Set[dto.OtpDTO](key, val, os.Cfg.Otp.ExpireTime, os.Redis)
	if err != nil {
		return &service_errors.ServiceError{EndUserMessage: "cant set otp try again later", Err: err}
	}
	return nil
}

func (os *OtpService) ValidateOTP(mobileNumber, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.DefaultRedisKey, mobileNumber)
	res, err := cache.Get[dto.OtpDTO](key, os.Redis)
	if err != nil {
		return err
	}
	if res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpInvalid}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set[dto.OtpDTO](key, *res, os.Cfg.Otp.ExpireTime, os.Redis)
		if err != nil {
			return err
		}
	}
	return nil

}
