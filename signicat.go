package signicat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// A Client manages communication with the Signicat API.
type Client struct {
	client  *http.Client
	baseURL *url.URL

	common service

	Signature *SignatureService
}

type service struct {
	client *Client
}

// NewClient returns a new Signicat API client. To use API methods which require authentication, provide an http.Client that will
// perform the authentication for you. Most likelely you want to use Oauth2 and the golang.org/x/oauth2 package.
func NewClient(httpClient *http.Client, baseURL string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  httpClient,
		baseURL: u,
	}

	c.common.client = c
	c.Signature = (*SignatureService)(&c.common)

	return c, nil
}

// NewRequest creates a new API request with the provided http method, body and with path which is the clients base url + relativeUrl.
func (c *Client) NewRequest(method, relativeURL string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(relativeURL)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do sends an API request. The response is decoded and stored in the value pointed to by v unless an error is returned.
func (c *Client) Do(req *http.Request, v interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// TODO: Better error handling. Make a custom error type with info?
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("received response with http code: %d", res.StatusCode)
	}

	if v != nil {
		// ignore EOF errors caused by empty response body
		if err = json.NewDecoder(res.Body).Decode(v); err != nil && err != io.EOF {
			return err
		}
	}

	return nil
}
