package http

import (
	"context"
	"net/http"
)

func (c *ClientWrapper) Post(
	ctx context.Context, endpoint string, params map[string]string, headers map[string][]string, body []byte,
) ([]byte, error) {
	return c.handleResponse(c.PostWithResp(ctx, endpoint, params, headers, body))
}

func (c *ClientWrapper) PostWithResp(
	ctx context.Context, endpoint string, params map[string]string, headers map[string][]string, body []byte,
) (*http.Response, error) {
	return c.WithResp(ctx, endpoint, http.MethodPost, params, headers, body)
}
