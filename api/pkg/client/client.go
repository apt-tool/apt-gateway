package client

import (
	"fmt"
	"net/http"
	"strings"
)

type HTTPClient interface {
	Get(uri string, headers ...string) (*http.Response, error)
}

// NewClient
// creating a new http client.
func NewClient() HTTPClient {
	return &client{
		conn: &http.Client{},
	}
}

type client struct {
	conn *http.Client
}

// Get
// making a get request.
func (c client) Get(uri string, headers ...string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("[client.Get] failed in creating requests: %w", err)
	}

	for _, pair := range headers {
		parts := strings.Split(pair, ":")

		req.Header.Add(parts[0], parts[1])
	}

	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[client.Get] http request failed: %w", err)
	}

	return resp, nil
}
