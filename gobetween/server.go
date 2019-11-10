package gobetween

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/thetechnick/gobetween-client-go/gobetween/api"
)

// ListServers returns all Servers that are registered in gobetween.
func (c *Client) ListServers(ctx context.Context) (map[string]api.Server, error) {
	req, err := c.request(ctx, http.MethodGet, "/servers", nil)
	if err != nil {
		return nil, err
	}

	res := map[string]api.Server{}
	if err = c.do(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetServer returns a server by it's name.
func (c *Client) GetServer(ctx context.Context, name string) (*api.Server, error) {
	req, err := c.request(ctx, http.MethodGet, "/servers/"+name, nil)
	if err != nil {
		return nil, err
	}

	res := &api.Server{}
	if err = c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateServer creates the server with the given name, if it does not exist.
func (c *Client) CreateServer(ctx context.Context, name string, cfg *api.Server) error {
	bodyData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	req, err := c.request(ctx, http.MethodPost, "/servers/"+name, bytes.NewReader(bodyData))
	if err != nil {
		return err
	}

	if err = c.do(req, nil); err != nil {
		return err
	}
	return nil
}

// Deletes the given server from gobetween.
func (c *Client) DeleteServer(ctx context.Context, name string) error {
	req, err := c.request(ctx, http.MethodDelete, "/servers/"+name, nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
