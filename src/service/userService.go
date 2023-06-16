package service

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/utils"
	"github.com/sirupsen/logrus"
)

type UserService struct{}

func (userService *UserService) RegisterUser(request dto.RegisterUserRequest) (*dto.RegisterUserResponse, *dto.Error) {
	logrus.Info("UserService.RegisterUser")

	// Preprocessing
	request.Email = strings.ToLower(request.Email)

	// Validation
	var fieldErrors []string
	if request.Name == "" {
		fieldErrors = append(fieldErrors, "name should not be empty")
	}

	if request.Email == "" {
		fieldErrors = append(fieldErrors, "email should not be empty")
	}

	if request.Password == "" {
		fieldErrors = append(fieldErrors, "password should not be empty")
	} else if len(request.Password) < 6 {
		fieldErrors = append(fieldErrors, "password should be at least 6 characters")
	}

	if request.Phone == "" {
		fieldErrors = append(fieldErrors, "phone should not be empty")
	} else if len(request.Phone) < 10 {
		fieldErrors = append(fieldErrors, "phone must contain at least 10 digits")
	} else {
		_, err := strconv.Atoi(request.Phone)
		if err != nil {
			fieldErrors = append(fieldErrors, "phone must contain only numbers")
		}
	}

	// TODO: Check if email already exists

	if len(fieldErrors) > 0 {
		return nil, &dto.Error{
			Code:    http.StatusBadRequest,
			Message: fieldErrors,
		}
	}

	// Hash Password
	hash, err := utils.HashPassword(request.Password)
	if err != nil {
		logrus.Error("Failed to Hash Password", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to Register User",
		}
	}

	request.Password = hash
	user, err := userRepository.RegisterUser(request)
	if err != nil {
		logrus.Error("Failed to Register User", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to Register User",
		}
	}

	return &dto.RegisterUserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}

func (userService *UserService) LoginUser(request dto.LoginUserRequest) (*dto.LoginUserResponse, *dto.Error) {
	logrus.Info("userService.LoginUser")
	request.Email = strings.ToLower(request.Email)
	// Get user using email.
	user, err := userRepository.GetUserByEmail(request.Email)
	if err != nil {
		logrus.Error("Failed to get User", err)
		return nil, &dto.Error{
			Code:    http.StatusForbidden,
			Message: "Invalid Credentials",
		}
	}

	// Validate Password
	if !utils.CheckPasswordHash(request.Password, user.Password) {
		logrus.Error("userService.LoginUser: Invalid Password for user: " + request.Email)
		return nil, &dto.Error{
			Code:    http.StatusForbidden,
			Message: "Invalid Credentials",
		}
	}

	// Generate JWT
	token, err := utils.GenerateJwt(map[string]interface{}{
		"userId": strconv.Itoa(user.Id),
		"role":   user.Role,
		"email":  user.Email,
	})

	if err != nil {
		logrus.Error("Faild to generate Token", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Faild to Authenticate",
		}
	}

	return &dto.LoginUserResponse{
		Id:            user.Id,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
		PhoneVerified: user.PhoneVerified,
		ApplicationId: user.ApplicationId,
		JwtToken:      token,
	}, nil
}

func (userService *UserService) GetUser(userId int) (*dto.LoginUserResponse, *dto.Error) {
	logrus.Info("userService.GetUser")
	user, err := userRepository.GetUserById(strconv.Itoa(userId))
	if err != nil {
		logrus.Error("Failed to get User", err)
		return nil, &dto.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Credentials",
		}
	}

	token, err := utils.GenerateJwt(map[string]interface{}{
		"userId": strconv.Itoa(user.Id),
		"role":   user.Role,
		"email":  user.Email,
	})

	if err != nil {
		logrus.Error("Faild to generate Token", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Faild to Authenticate",
		}
	}

	return &dto.LoginUserResponse{
		Id:            user.Id,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
		PhoneVerified: user.PhoneVerified,
		ApplicationId: user.ApplicationId,
		JwtToken:      token,
	}, nil
}
