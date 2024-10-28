package usecases_test

import (
	"context"
	"errors"
	mock_repository "github.com/volkankocaali/bi-taksi-case/internal/repository/mocks"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/request"
	"github.com/volkankocaali/bi-taksi-case/internal/usecases"
)

func TestRegister(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockUserRepository(ctrl)
		uc := usecases.NewUserUseCase(mockRepo)

		user := request.UserRegisterSchema{
			Username:             "testuser",
			Password:             "password123",
			PasswordConfirmation: "password123",
		}

		mockRepo.EXPECT().CheckUserExists(gomock.Any(), user.Username).Return(false, nil)
		mockRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)

		err := uc.Register(context.Background(), user, "test-issuer")
		assert.NoError(t, err)
	})

	t.Run("PasswordMismatch", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockUserRepository(ctrl)
		uc := usecases.NewUserUseCase(mockRepo)

		user := request.UserRegisterSchema{
			Username:             "testuser",
			Password:             "password123",
			PasswordConfirmation: "password124",
		}

		err := uc.Register(context.Background(), user, "test-issuer")
		assert.EqualError(t, err, "password and confirmation password do not match")
	})
}

func TestRegister_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	uc := usecases.NewUserUseCase(mockRepo)

	user := request.UserRegisterSchema{
		Username:             "testuser",
		Password:             "password123",
		PasswordConfirmation: "password123",
	}

	mockRepo.EXPECT().CheckUserExists(gomock.Any(), user.Username).Return(true, nil)

	err := uc.Register(context.Background(), user, "test-issuer")
	assert.EqualError(t, err, "user already exists")
}

func TestRegister_CheckUserExistsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	uc := usecases.NewUserUseCase(mockRepo)

	user := request.UserRegisterSchema{
		Username:             "testuser",
		Password:             "password123",
		PasswordConfirmation: "password123",
	}

	mockRepo.EXPECT().CheckUserExists(gomock.Any(), user.Username).Return(false, errors.New("db error"))

	err := uc.Register(context.Background(), user, "test-issuer")
	assert.EqualError(t, err, "db error")
}

func TestRegister_RegisterError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	uc := usecases.NewUserUseCase(mockRepo)

	user := request.UserRegisterSchema{
		Username:             "testuser",
		Password:             "password123",
		PasswordConfirmation: "password123",
	}

	mockRepo.EXPECT().CheckUserExists(gomock.Any(), user.Username).Return(false, nil)
	mockRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(errors.New("db error"))

	err := uc.Register(context.Background(), user, "test-issuer")
	assert.EqualError(t, err, "db error")
}
