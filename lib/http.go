package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HTTPClient struct {
	method string
	url    string
	body   []byte
	token  string
}

func NewHttpRequest(method string, url string, body []byte) *HTTPClient {
	return &HTTPClient{
		method: method,
		url:    url,
		body:   body,
	}
}

func (h *HTTPClient) WithToken(token string) *HTTPClient {
	h.token = token
	return h
}

func (h *HTTPClient) Download(filename string) error {
	req, err := http.NewRequest(h.method, h.url, nil)
	if err != nil {
		return fmt.Errorf("failed to new request: %v", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to invoke API: %v", err)
	}

	defer resp.Body.Close()
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create video file: %v", err)
	}
	defer out.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to store video: %v", err)
	}
	return nil
}

func (h *HTTPClient) Run(v any) error {
	var reqBody io.Reader = nil
	if h.body != nil {
		reqBody = bytes.NewReader(h.body)
	}
	req, err := http.NewRequest(h.method, h.url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to new request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	if h.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.token))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to invoke API: %v", err)
	}
	body := resp.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read data: %v", err)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to unmarshal data: %v", err)
	}
	return nil
}
