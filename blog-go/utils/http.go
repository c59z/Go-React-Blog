package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// HttpRequest sends an HTTP request with optional headers, query params, and body.
func HttpRequest(
	urlStr string, // request URL
	method string, // HTTP method (GET, POST, etc.)
	headers map[string]string, // headers
	params map[string]string, // query parameters
	data any, // request body (optional)
) (*http.Response, error) {

	// Parse URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	query := u.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	u.RawQuery = query.Encode()

	// Encode body as JSON if provided
	buf := new(bytes.Buffer)
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	// Create HTTP request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Default Content-Type for body
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
