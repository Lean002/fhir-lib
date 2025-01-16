package client

import (
	"net/http"
)

// FHIRClient FHIR client structure
type FHIRClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Username   string
	Password   string
}

// NewFHIRClient creates a new FHIR client with basic authentication
func NewFHIRClient(baseURL, username, password string) *FHIRClient {
	return &FHIRClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
		Username:   username,
		Password:   password,
	}
}
