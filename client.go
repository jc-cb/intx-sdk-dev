package intx

import (
	"net/http"
)

// var defaultV1ApiBaseUrl = "https://api.international.coinbase.com/api/v1"
var defaultV1ApiBaseUrl = "https://api-n5e1.coinbase.com/api/v1"

type Client struct {
	HttpClient  http.Client
	Credentials *Credentials
	HttpBaseUrl string
}

func (c *Client) BaseUrl(u string) *Client {
	c.HttpBaseUrl = u
	return c
}

func NewClient(credentials *Credentials, httpClient http.Client) *Client {
	return &Client{
		Credentials: credentials,
		HttpClient:  httpClient,
		HttpBaseUrl: defaultV1ApiBaseUrl,
	}
}
