package circuitbreaker

import (
	"bytes"
	"fmt"
	"github.com/sony/gobreaker/v2"
	"io"
	"net/http"
	"time"
)

var cb *gobreaker.CircuitBreaker[[]byte]

func InitCircuitBreaker() {
	st := gobreaker.Settings{
		Name:        "DriverLocationService",
		MaxRequests: 1,
		Interval:    60 * time.Second,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
	}

	cb = gobreaker.NewCircuitBreaker[[]byte](st)

}

func Post(url string, data []byte, apiKey *string) ([]byte, error) {
	return requestWithCircuitBreaker(http.MethodPost, url, data, apiKey)
}

func requestWithCircuitBreaker(method, url string, data []byte, token *string) ([]byte, error) {
	body, err := cb.Execute(func() ([]byte, error) {
		var reqBody io.Reader
		if data != nil {
			reqBody = bytes.NewBuffer(data)
		}

		req, err := http.NewRequest(method, url, reqBody)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		if token != nil {
			req.Header.Set("Authorization", "Bearer "+*token)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})

	if err != nil {
		return nil, err
	}

	return body, nil
}
