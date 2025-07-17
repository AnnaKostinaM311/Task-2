package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"project/config"
	"project/internal/interfaces"
)

type predictionService struct {
	client interfaces.HTTPClient
	cfg    *config.Config
}

func NewPredictionService(client interfaces.HTTPClient, cfg *config.Config) interfaces.PredictionService {
	return &predictionService{
		client: client,
		cfg:    cfg,
	}
}

func (s *predictionService) Predict(ctx context.Context, endpoint string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("json marshal error: %w", err)
	}

	url := s.cfg.API.BaseURL + "/" + endpoint
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("request creation error: %w", err)
	}
	req.Header.Set("Authorization", s.cfg.Server.AuthTokens[0])
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("api error: %s", string(body))
	}

	return io.ReadAll(resp.Body)
}
