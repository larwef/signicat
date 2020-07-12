package signicat

import (
	"net/http"
)

type Client struct {
	client *http.Client
}

func New(httpClient *http.Client) *Client {
	return &Client{
		client: httpClient,
	}
}
