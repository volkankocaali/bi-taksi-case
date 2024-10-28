package handler

import (
	"encoding/json"
	"github.com/volkankocaali/bi-taksi-case/config"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/request"
	"github.com/volkankocaali/bi-taksi-case/internal/usecases"
	"github.com/volkankocaali/bi-taksi-case/pkg/utils"
	"net/http"
)

type UserHandler struct {
	usecase *usecases.UserUseCase
}

func NewUserHandler(usecase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var user request.UserLoginSchema
	validator := utils.NewValidator()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusBadRequest,
				Message: "failed to decode request body",
			})
		return
	}

	errors := validator.ValidateStruct(&user)
	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"errors": errors})
		return
	}

	secretKey := config.GetConfig().JwtSecretKey
	issuer := config.GetConfig().JwtIssuer
	userResponse, err := h.usecase.Login(r.Context(), user.Username, user.Password, secretKey, issuer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

func (h *UserHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var user request.UserRegisterSchema
	validator := utils.NewValidator()

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusBadRequest,
				Message: "failed to decode request body",
			})
		return
	}

	errors := validator.ValidateStruct(&user)
	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"errors": errors})
		return
	}

	secretKey := config.GetConfig().JwtSecretKey
	issuer := config.GetConfig().JwtIssuer

	// user register
	err := h.usecase.Register(r.Context(), user, issuer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// user login
	login, err := h.usecase.Login(r.Context(), user.Username, user.Password, secretKey, issuer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusBadRequest,
				Message: "user not login",
			})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(login)
}
