package interfaces

import (
	"context"
	"net/http"
)

type PredictionService interface {
	Predict(ctx context.Context, endpoint string, data interface{}) ([]byte, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
