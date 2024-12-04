package http

import (
	"context"
	"net/http"
)

func (c *ClientWrapper) Get(
	ctx context.Context, endpoint string, headers map[string][]string, params map[string]string,
) ([]byte, error) {
	return c.handleResponse(c.GetWithResp(ctx, endpoint, headers, params))
}

func (c *ClientWrapper) GetWithResp(
	ctx context.Context, endpoint string, headers map[string][]string, params map[string]string,
) (*http.Response, error) {
	return c.WithResp(ctx, endpoint, http.MethodGet, params, headers, nil)
}
