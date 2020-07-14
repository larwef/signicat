package signicat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.idfy.io/"
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

// NewClient returns a new client with the default base url.
func NewClient(httpClient *http.Client) *Client {
	client, err := NewClientWithURL(httpClient, defaultBaseURL)
	if err != nil {
		panic("unable to initiate default client. This should not happen")
	}

	return client
}

// NewClientWithURL returns a new Signicat API client. To use API methods which require authentication, provide an http.Client
// that will perform the authentication for you. Most likelely you want to use Oauth2 and the golang.org/x/oauth2 package.
func NewClientWithURL(httpClient *http.Client, baseURL string) (*Client, error) {
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
		// Write directly if v implements io.Writer.
		if w, ok := v.(io.Writer); ok {
			if _, err := io.Copy(w, res.Body); err != nil {
				return err
			}
		} else {
			// Ignore EOF errors caused by empty response body.
			if err = json.NewDecoder(res.Body).Decode(v); err != nil && err != io.EOF {
				return err
			}
		}
	}

	return nil
}
