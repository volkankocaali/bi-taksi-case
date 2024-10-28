package handler

import (
	"encoding/json"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/request"
	"github.com/volkankocaali/bi-taksi-case/internal/usecases"
	"github.com/volkankocaali/bi-taksi-case/pkg/utils"
	"net/http"
)

// MatchingHandler handles driver matching operations
type MatchingHandler struct {
	usecase *usecases.MatchingUseCase
}

// NewMatchingHandler creates a new MatchingHandler
// @Summary Initialize MatchingHandler
// @Description Initializes a new matching handler with the provided use case
func NewMatchingHandler(usecase *usecases.MatchingUseCase) *MatchingHandler {
	return &MatchingHandler{
		usecase: usecase,
	}
}

// MatchDriver finds the nearest driver within a specified radius
// @Summary Find nearest driver
// @Description Finds the nearest driver to a specified latitude and longitude within a given radius
// @Tags Matching
// @Accept json
// @Produce json
// @Param matchRequest body MatchRequest true "Matching parameters"
// @Success 200 {object} response.Driver "Matched driver location"
// @Failure 400 {object} handler.ApiResponse "Bad request - validation error"
// @Failure 500 {object} handler.ApiResponse "Internal server error"
// @Router /match-driver [post]
func (h *MatchingHandler) MatchDriver(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body into MatchRequest struct
	var req request.MatchRequest
	validator := utils.NewValidator()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request payload",
		})
		return
	}

	errors := validator.ValidateStruct(&req)
	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"errors": errors})
		return
	}

	driver, err := h.usecase.FindNearestDriver(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(driver)
}
