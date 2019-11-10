package gobetween

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Config holds Client configuration options.
type Config struct {
	// Http user agent to send with every request.
	UserAgent string

	// Auth holds Authentication information for each request.
	Auth auth
}

// auth adds authentication information to a request.
type auth interface {
	Apply(req *http.Request)
}

// ClientOption configures settings on the client.
type ClientOption func(*Config)

// WithUserAgent sets the user agent on the client.
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Config) {
		c.UserAgent = userAgent
	}
}

// WithAuth sets per request authentication on the client.
func WithAuth(a auth) ClientOption {
	return func(c *Config) {
		c.Auth = a
	}
}

// BasicAuth
type BasicAuth struct {
	Username, Password string
}

// Apply adds the basic auth credentials to the request.
func (a *BasicAuth) Apply(req *http.Request) {
	req.SetBasicAuth(a.Username, a.Password)
}

type Client struct {
	// Endpoint of the gobetween API in format: http/https://host:port
	endpoint string

	config *Config
	client *http.Client
}

// NewClient creates a new client for the given endpoint and options.
func NewClient(endpoint string, opts ...ClientOption) *Client {
	c := &Client{
		endpoint: endpoint,
		config:   &Config{},
		client:   &http.Client{},
	}
	for _, opt := range opts {
		opt(c.config)
	}
	return c
}

func (c *Client) request(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	url := c.endpoint + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %v", err)
	}
	req.Header.Set("User-Agent", c.config.UserAgent)
	if c.config.Auth != nil {
		c.config.Auth.Apply(req)
	}
	req = req.WithContext(ctx)
	return req, nil
}

func (c *Client) do(r *http.Request, v interface{}) error {
	resp, err := c.client.Do(r)
	if err != nil {
		return fmt.Errorf("sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}

	// errors
	switch resp.StatusCode {
	case 200:
		// no error
	case 400:
		return newBadRequestError(body)
	case 401:
		return newUnauthorizedError(body)
	case 404:
		return newNotFoundError(body)
	case 409:
		return newConflictError(body)
	case 500:
		return newInternalError(body)
	default:
		return fmt.Errorf("unknown error: (%d) %s", resp.StatusCode, string(body))
	}

	if v != nil {
		return json.Unmarshal(body, v)
	}

	return nil
}
