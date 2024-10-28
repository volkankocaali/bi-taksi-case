package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/volkankocaali/bi-taksi-case/internal/auth"
	"github.com/volkankocaali/bi-taksi-case/internal/domain"
	"github.com/volkankocaali/bi-taksi-case/internal/repository"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/request"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/response"
	"github.com/volkankocaali/bi-taksi-case/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseInterface interface {
	Register(ctx context.Context, user request.UserRegisterSchema) error
	Login(ctx context.Context, username, password string) (response.UserResponse, error)
}

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, user request.UserRegisterSchema, issuer string) error {
	// confirmation password check
	if user.Password != user.PasswordConfirmation {
		return errors.New("password and confirmation password do not match")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// check if user already exists
	if exists, checkErr := uc.userRepo.CheckUserExists(ctx, user.Username); checkErr != nil {
		return checkErr
	} else if exists {
		return errors.New("user already exists")
	}

	// register user
	return uc.userRepo.Register(ctx, domain.User{
		Username: user.Username,
		Password: user.Password,
		Service:  issuer,
	})
}

func (uc *UserUseCase) Login(ctx context.Context, username, password, secretKey, issuer string) (response.UserResponse, error) {
	user, err := uc.userRepo.Login(ctx, username)
	if err != nil {
		return response.UserResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return response.UserResponse{}, errors.New("password is incorrect")
	}

	// check issuer
	switch issuer {
	case "driver-service":
		if issuer != user.Service {
			logger.Error(fmt.Sprintf("user %s is not registered to the service. Service name : %s", user.Username, issuer))
			return response.UserResponse{}, errors.New("user is not registered to the service")
		}
	case "matching-service":
		break
	}

	// generate token
	token, err := auth.GenerateApiJWT(secretKey, issuer, user.ID, user.Username)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to generate token for user %s", user.Username))
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		Status: 200,
		Token:  token,
		User: response.UserDTO{
			ID:       user.ID,
			Username: user.Username,
		},
		Message: "User logged in successfully",
	}, nil
}
