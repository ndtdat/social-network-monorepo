package http

import (
	"fmt"
	"io"
	"net/http"
)

func (c *ClientWrapper) handleResponse(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %s, body: %s", resp.Status, string(body))
	}

	return body, nil
}
