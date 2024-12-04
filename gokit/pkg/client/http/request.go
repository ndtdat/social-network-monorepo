package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"io"
	"net/http"
	"time"
)

func (c *ClientWrapper) createRequest(
	ctx context.Context, endpoint, method string, params map[string]string, headers map[string][]string, body []byte,
) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)
	if body != nil {
		req, err = http.NewRequestWithContext(ctx, method, endpoint, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequestWithContext(ctx, method, endpoint, nil)
	}

	if err != nil {
		return nil, err
	}

	// Set header
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	for k, values := range headers {
		for _, value := range values {
			req.Header.Set(k, value)
		}
	}

	if params != nil {
		// Set request params
		query := req.URL.Query()
		for key, value := range params {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}

	return req, err
}

func (c *ClientWrapper) WithResp(
	ctx context.Context, endpoint, method string, params map[string]string, headers map[string][]string, body []byte,
) (*http.Response, error) {
	request, err := c.createRequest(ctx, endpoint, method, params, headers, body)
	if err != nil {
		return nil, err
	}

	return c.HTTPClient.Do(request)
}

func (c *ClientWrapper) UnpackRespMap(
	ctx context.Context, endpoint, method string, params map[string]string, headers map[string][]string, payload any,
	maxAttempt int, sleepBetweenAttempt *time.Duration, resultObject any,
) error {
	var (
		body []byte
		err  error
	)

	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return nil
		}
	}

	var resp *http.Response
	results, err := util.Retry(func() (util.DoResult, error) {
		//nolint:bodyclose
		reply, err := c.WithResp(ctx, endpoint, method, params, headers, body)
		if err != nil {
			return nil, err
		}

		statusCode := reply.StatusCode
		isSuccessful := statusCode == http.StatusOK
		if !isSuccessful {
			if reply != nil {
				defer reply.Body.Close()
			}

			bodyBytes, err := io.ReadAll(reply.Body)
			if err != nil {
				return nil, fmt.Errorf("status code: %v, cannot read body due to: %v", statusCode, err)
			}

			return nil, fmt.Errorf("status code: %v, msg: %s", statusCode, bodyBytes)
		}

		return util.DoResult{reply}, nil
	}, maxAttempt, sleepBetweenAttempt)
	if err != nil {
		return err
	}
	resp = results[0].(*http.Response)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resultObject == nil {
		return nil
	}

	respBody := resp.Body
	bodyBytes, err := io.ReadAll(respBody)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bodyBytes, &resultObject); err != nil {
		return fmt.Errorf("cannot unmarshal response %s due to %v", string(bodyBytes), err)
	}

	return nil
}
