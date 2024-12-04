package http

import (
	"net/http"

	"go.uber.org/zap"
)

type ClientWrapper struct {
	logger     *zap.Logger
	HTTPClient *http.Client
}

func NewClientWrapper(logger *zap.Logger) *ClientWrapper {
	t := http.DefaultTransport.(*http.Transport).Clone() //nolint:forcetypeassert
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient := &http.Client{
		Transport: t,
	}

	return &ClientWrapper{
		logger:     logger,
		HTTPClient: httpClient,
	}
}
