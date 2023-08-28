package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	URL    string
	Client *http.Client
}

func New(url string) *Client {
	return &Client{
		URL:    url,
		Client: &http.Client{},
	}
}

func (c *Client) Query(name string, value string) string {
	return name + "=" + value
}

func (c *Client) Request(method string, authHeader string, route string, body io.Reader, queries ...string) ([]byte, error) {
	path := c.URL + route
	log.Println(path)

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+authHeader)

	q := req.URL.Query()
	for _, query := range queries {
		sp := strings.SplitN(query, "=", 2)
		if len(sp) != 2 {
			continue
		}
		q.Add(sp[0], sp[1])
	}
	req.URL.RawQuery = q.Encode()
	log.Printf("Request: %s %s\n", method, req.URL.String())

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Status: %d, Msg: %s\n", resp.StatusCode, string(data))
	}

	return data, nil
}

func (c *Client) Get(route string, authHeader string, queries ...string) ([]byte, error) {
	return c.Request(http.MethodGet, authHeader, route, nil, queries...)
}

func (c *Client) Post(route string, authHeader string, body io.Reader, queries ...string) ([]byte, error) {
	return c.Request(http.MethodPost, authHeader, route, body, queries...)
}

func (c *Client) Put(route string, authHeader string, body io.Reader, queries ...string) ([]byte, error) {
	return c.Request(http.MethodPut, authHeader, route, body, queries...)
}

func (c *Client) Patch(route string, authHeader string, body io.Reader, queries ...string) ([]byte, error) {
	return c.Request(http.MethodPatch, authHeader, route, body, queries...)
}

func (c *Client) Delete(route string, authHeader string, body io.Reader, queries ...string) ([]byte, error) {
	return c.Request(http.MethodDelete, authHeader, route, body, queries...)
}
