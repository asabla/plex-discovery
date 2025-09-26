package plex

import (
	"context"
	"net/http"

	"github.com/asabla/plex-discovery/internal/plex/clients"
)

type Client struct {
	transport *clients.ClientWithResponses
}

func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c, err := clients.NewClientWithResponses(baseURL, clients.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	return &Client{transport: c}, nil
}

