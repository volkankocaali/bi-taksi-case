package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/request"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/response"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DriverLocation struct {
	DriverID  string    `json:"driver_id"`
	Location  GeoJSON   `json:"location"`
	Timestamp time.Time `json:"timestamp"`
}

type GeoJSON struct {
	Type        string     `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

func main() {
	file, err := os.Open("cmd/import-data/driver_locations.csv")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var driverLocation []DriverLocation
	for key, record := range records {
		if key == 0 {
			continue
		}

		lat, _ := strconv.ParseFloat(record[0], 64)
		lng, _ := strconv.ParseFloat(record[1], 64)

		driverLocation = append(driverLocation, DriverLocation{
			DriverID: generateDriverID(),
			Location: GeoJSON{
				Type: "Point",
				Coordinates: [2]float64{
					lat,
					lng,
				},
			},
			Timestamp: time.Now(),
		})
	}

	token, err := loginToApi(request.UserLoginSchema{
		Username: "john.doe",
		Password: "bitaksi",
	})
	sendToAPI(driverLocation, token)
}

func loginToApi(userRequest request.UserLoginSchema) (string, error) {
	url := "http://localhost:8081/api/v1/login"
	jsonData, err := json.Marshal(userRequest)

	if err != nil {
		fmt.Println("Error marshalling data", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return "", err
	}

	defer resp.Body.Close()

	var loginResponse response.UserResponse
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		fmt.Println("Error decoding response", err)
		return "", err
	}

	return loginResponse.Token, nil
}

func sendToAPI(driverLocation []DriverLocation, token string) {
	url := "http://localhost:8081/api/v1/driver-locations"
	jsonData, err := json.Marshal(driverLocation)

	if err != nil {
		fmt.Println("Error marshalling data", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func generateDriverID() string {
	return "driver-" + strconv.Itoa(rand.Intn(1000))
}
