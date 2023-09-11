package services

import (
	"database/sql"
	"fmt"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/common"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Logger logging.Logger
	Cfg    *config.Config
	Otp    *OtpService
	Token  *TokenService
	Db     *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	otp := NewOTPService(cfg)
	token := NewTokenService(cfg)
	db := db.GetDB()
	return &UserService{
		Logger: logger,
		Cfg:    cfg,
		Otp:    otp,
		Token:  token,
		Db:     db,
	}
}

func (us *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := us.Otp.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	err := us.Db.Model(&models.User{}).Select("count(*) > 1").Where("email = ?", email).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (us *UserService) existsByPhone(phone string) (bool, error) {
	var exists bool
	err := us.Db.Model(&models.User{}).Select("count(*) > 0").Where("phone_number = ?", phone).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}
func (us *UserService) existsByUserName(username string) (bool, error) {
	var exists bool
	err := us.Db.Model(&models.User{}).Select("count(*) > 1").Where("username = ?", username).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (us *UserService) getDefaultRole() (roleId int, err error) {
	if err := us.Db.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).
		Error; err != nil {
		return -9, err
	}
	return roleId, nil
}

func (us *UserService) RegisterByUsername(req *dto.RegisterUserByUsername) error {
	user := &models.User{
		FirstName: req.FirstName,
		LastName:  sql.NullString{Valid: true, String: req.LastName},
		Username:  req.Username,
		Email:     sql.NullString{Valid: true, String: req.Email},
		Enable:    true,
	}
	exists, err := us.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = us.existsByUserName(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	bs, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bs)
	roleId, err := us.getDefaultRole()
	if err != nil {
		return err
	}
	tx := us.Db.Begin()
	err = tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		us.Logger.Error(err, logging.Postgres, logging.Insert, "cant add user", nil)
		return err
	}
	ur := &models.UserRole{UserId: user.ID, RoleId: roleId}
	err = tx.Create(ur).Error
	if err != nil {
		tx.Rollback()
		us.Logger.Error(err, logging.Postgres, logging.Insert, "cant add user", nil)
		return err
	}
	tx.Commit()

	return nil
}

func (us *UserService) RegisterLoginByPhone(req *dto.RegisterLoginByPhone) (*dto.TokenDetail, error) {
	err := us.Otp.ValidateOTP(req.Phone, req.Otp)
	if err != nil {
		return nil, err
	}
	exists, err := us.existsByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	var user = models.User{Username: req.Phone, PhoneNumber: sql.NullString{Valid: true, String: req.Phone}}
	if exists {
		var usr models.User
		err = us.Db.Model(&models.User{}).Where("username = ?", req.Phone).Preload("UserRoles.Role").Find(&usr).Error
		if err != nil {
			return nil, err
		}
		tDTO := dto.TokenDTO{
			Username: usr.Username,
			UserId:   usr.ID,
			FullName: fmt.Sprintf("%s %s", usr.FirstName, usr.LastName.String),
			Phone:    usr.PhoneNumber.String,
			Email:    usr.Email.String,
		}
		if len(usr.UserRoles) > 0 {
			for _, r := range usr.UserRoles {
				tDTO.Roles = append(tDTO.Roles, r.Role.Name)
			}
		}

		tk, err := us.Token.GenerateToken(&tDTO)
		if err != nil {
			return nil, err
		}
		return tk, nil
	} else {
		bs, err := bcrypt.GenerateFromPassword([]byte("changeMe"), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(bs)
		roleId, err := us.getDefaultRole()
		if err != nil {
			return nil, err
		}
		tx := us.Db.Begin()
		err = tx.Create(&user).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		ur := &models.UserRole{UserId: user.ID, RoleId: roleId}
		err = tx.Create(ur).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()

		var createdUser models.User

		err = us.Db.Model(&models.User{}).Where("username = ?", user.Username).Preload("UserRoles.Role").Find(&createdUser).Error
		if err != nil {
			return nil, err
		}
		tDTO := &dto.TokenDTO{
			UserId:   createdUser.ID,
			FullName: fmt.Sprintf("%s %s", createdUser.FirstName, createdUser.LastName.String),
			Username: createdUser.Username,
			Phone:    createdUser.PhoneNumber.String,
			Email:    createdUser.Email.String,
		}
		if len(createdUser.UserRoles) > 0 {
			for _, r := range createdUser.UserRoles {
				tDTO.Roles = append(tDTO.Roles, r.Role.Name)
			}
		}
		tk, err := us.Token.GenerateToken(tDTO)
		if err != nil {
			return nil, err
		}
		return tk, nil
	}
}

func (us *UserService) LoginByUsername(req *dto.LoginByUsername) (*dto.TokenDetail, error) {
	var user models.User
	err := us.Db.Model(&models.User{}).Where("username = ?", req.Username).Preload("UserRoles.Role").Find(&user).Error
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
	tk, err := us.Token.GenerateToken(tDTO)
	if err != nil {
		return nil, err
	}
	return tk, nil
}
