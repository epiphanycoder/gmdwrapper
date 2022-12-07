package net

import (
	"net/http"
	"time"
)

const defaultTimeoutSeconds = 15

// NewHTTPClient creates a HTPP client with a 15seconds.
func NewHTTPClient() *http.Client {
	return &http.Client{Timeout: time.Second * defaultTimeoutSeconds}
}

// SetUserAgent sets User-Agent to HTTP request header
func SetUserAgent(req *http.Request, userAgent string) {
	req.Header.Set("User-Agent", userAgent)
}

// SetHeader sets header and corresponding value to HTTP request header
func SetHeader(req *http.Request, header, value string) {
	req.Header.Set(header, value)
}
