package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"net/http"
)

type Client struct {
	client     *http.Client
	apiVersion string
}

var (
	ErrNoSuchImage = errors.New("no such image")
)

const (
	DockerNetWork = "unix"
	DockerAddress = "/var/run/docker.sock"
)

func NewClient(network, address string) (*Client, error) {
	c := &Client{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial(network, address)
				},
			},
		},
	}

	v, err := c.Version()
	if err != nil {
		return nil, err
	}
	c.apiVersion = v

	return c, nil
}

func (c *Client) Version() (string, error) {
	resp, err := c.do("GET", "http://localhost/version", nil)
	if err != nil {
		return "", err
	}
	if resp != nil {
		defer func() {
			resp.Body.Close()
		}()
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received unexpected status %d while trying to retrieve the server version", resp.StatusCode)
	}

	var versionResponse map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&versionResponse); err != nil {
		return "", err
	}
	if version, ok := (versionResponse["ApiVersion"]).(string); ok {
		return version, nil
	}
	return "", nil
}

func (c *Client) do(method, url string, data any) (*http.Response, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(request)
}
