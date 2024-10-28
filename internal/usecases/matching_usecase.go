package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volkankocaali/bi-taksi-case/internal/repository"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/request"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/response"
	"github.com/volkankocaali/bi-taksi-case/pkg/circuitbreaker"
)

const DriverLocationServiceBaseUrl = "http://driver-location-api:8081/api/v1/"

const FindDriverWithinRadiusEndpoint = "find-driver-within-radius"
const Login = "login"

type MatchingUseCase struct {
	driverRepo repository.DriverLocationRepository
}

func NewMatchingUseCase(driverRepo repository.DriverLocationRepository) *MatchingUseCase {
	return &MatchingUseCase{
		driverRepo: driverRepo,
	}
}

func (u *MatchingUseCase) FindNearestDriver(req request.MatchRequest) (*response.Driver, error) {
	findDriverWithinRadiusUrl := fmt.Sprintf(DriverLocationServiceBaseUrl + FindDriverWithinRadiusEndpoint)

	findDriverWithinRadiusData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshal request: %v", err)
	}

	// first driver location service login

	var loginResponse response.UserResponse
	loginUrl := fmt.Sprintf(DriverLocationServiceBaseUrl + Login)
	loginData := request.UserLoginSchema{
		Username: "john.doe",
		Password: "bitaksi",
	}

	jsonLoginData, err := json.Marshal(loginData)
	if err != nil {
		return nil, fmt.Errorf("error marshal request: %v", err)
	}

	loginBody, loginErr := circuitbreaker.Post(loginUrl, jsonLoginData, nil)
	if loginErr != nil {
		return nil, errors.New("error not reach the driver location service")
	}

	if err = json.Unmarshal(loginBody, &loginResponse); err != nil {
		return nil, fmt.Errorf("error unmarshal user response driver location service: %v", err)
	}

	token := loginResponse.Token
	body, err := circuitbreaker.Post(findDriverWithinRadiusUrl, findDriverWithinRadiusData, &token)
	if err != nil {
		return nil, errors.New("error not reach the driver location service")
	}

	var drivers []response.Driver

	if err = json.Unmarshal(body, &drivers); err != nil {
		return nil, fmt.Errorf("error unmarshal driver response: %v", err)
	}

	if len(drivers) == 0 {
		return nil, fmt.Errorf("no driver found")
	}

	return &drivers[0], nil
}
