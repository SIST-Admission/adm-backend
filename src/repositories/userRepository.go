package repositories

import (
	"github.com/SIST-Admission/adm-backend/src/db"
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/sirupsen/logrus"
)

type UserRepository struct{}

func (repo *UserRepository) RegisterUser(payload dto.RegisterUserRequest) (*models.User, error) {
	logrus.Info("UserRepository.RegisterUser")
	db := db.GetInstance()

	user := models.User{
		Name:          payload.Name,
		Email:         payload.Email,
		Password:      payload.Password,
		EmailVerified: false,
		PhoneVerified: false,
		Role:          "STUDENT",
		Phone:         payload.Phone,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
