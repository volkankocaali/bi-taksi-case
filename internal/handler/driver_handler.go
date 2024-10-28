package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/volkankocaali/bi-taksi-case/internal/domain"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/request"
	"github.com/volkankocaali/bi-taksi-case/internal/usecases"
	"github.com/volkankocaali/bi-taksi-case/pkg/logger"
	"github.com/volkankocaali/bi-taksi-case/pkg/utils"
	"net/http"
)

// DriverHandler handles driver location operations
type DriverHandler struct {
	usecase *usecases.DriverLocationUseCase
}

// NewDriverHandler creates a new DriverHandler
// @Summary Initialize DriverHandler
// @Description Initializes a new driver handler with the provided use case
func NewDriverHandler(usecase *usecases.DriverLocationUseCase) *DriverHandler {
	return &DriverHandler{usecase: usecase}
}

// DriverCreateOrUpdate creates or updates driver locations
// @Summary Create or Update Driver Locations
// @Description Processes a batch of driver locations for creation or updating
// @Tags Driver
// @Accept json
// @Produce json
// @Param locations body []domain.DriverLocation true "Driver locations to process"
// @Success 200 {object} ApiResponse "Locations processed successfully"
// @Failure 400 {object} ApiResponse "Failed to decode request body"
// @Failure 500 {object} ApiResponse "Failed to process locations"
// @Router /driver-locations [post]
func (h *DriverHandler) DriverCreateOrUpdate(w http.ResponseWriter, r *http.Request) {
	var locations []domain.DriverLocation

	if err := json.NewDecoder(r.Body).Decode(&locations); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusBadRequest,
				Message: "failed to decode request body",
			})
		logger.Error("Invalid request body:", err)
		return
	}

	err := h.usecase.CreateOrUpdateDriverLocations(r.Context(), locations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusInternalServerError,
				Message: "failed to process locations",
			})
		logger.Error("Failed to process locations:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiResponse{
		Status:  http.StatusOK,
		Message: "locations processed successfully",
	})

}

// GetLatestDriverLocation retrieves the latest location of a driver
// @Summary Get Latest Driver Location
// @Description Retrieves the latest known location of the specified driver
// @Tags Driver
// @Produce json
// @Param driver_id path string true "Driver ID"
// @Success 200 {object} domain.DriverLocation "Latest driver location"
// @Failure 404 {object} ApiResponse "Driver location not found"
// @Failure 500 {object} ApiResponse "Internal server error"
// @Router /driver-locations/{driver_id} [get]
func (h *DriverHandler) GetLatestDriverLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	driverID := vars["driver_id"]

	ctx := r.Context()
	location, err := h.usecase.GetLatestDriverLocation(ctx, driverID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		logger.Error("Error get driver %s location &v", driverID, err)
		return
	}

	if location == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusNotFound,
				Message: "driver location not found",
			})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

// FindDriverWithinRadius finds drivers within a specified radius
// @Summary Find Drivers Within Radius
// @Description Finds drivers within a radius from a specified location
// @Tags Driver
// @Accept json
// @Produce json
// @Param request body request.DriverRequest true "Search parameters (latitude, longitude, radius)"
// @Success 200 {array} domain.DriverLocation "List of drivers within radius"
// @Failure 400 {object} ApiResponse "Invalid request payload or validation error"
// @Failure 500 {object} ApiResponse "Internal server error"
// @Router /find-driver-within-radius [post]
func (h *DriverHandler) FindDriverWithinRadius(w http.ResponseWriter, r *http.Request) {
	var req request.DriverRequest
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

	page := 1
	pageSize := 10
	if req.Page > 0 {
		page = req.Page
	}
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}

	drivers, err := h.usecase.FindDriversWithinRadius(r.Context(), req.Lat, req.Lon, req.Radius, page, pageSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			ApiResponse{
				Status:  http.StatusInternalServerError,
				Message: "internal server error",
			})
		logger.Error("Error get driver within radius %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}
