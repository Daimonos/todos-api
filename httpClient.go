package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Post is a helper for posting to an http endpoint
func Post(endpoint string, payload map[string]string) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return http.Post(endpoint, "application/json", bytes.NewBuffer(body))
}
