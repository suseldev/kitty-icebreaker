package firecracker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type client struct {
	http *http.Client
}

func newClient(socketPath string) *client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: func(
				ctx context.Context,
				_, _ string,
			) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(
					ctx,
					"unix",
					socketPath,
				)
			},
		},
	}

	return &client{
		http: httpClient,
	}
}
func (c *client) put(path string, body any) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost"+path, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("firecracker API returned %s: %s", resp.Status, data)
	}

	return nil
}
