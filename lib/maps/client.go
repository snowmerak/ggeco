package maps

import "googlemaps.github.io/maps"

// jetti:bean Client
type Client struct {
	baseClient *maps.Client
}

func New(apiKey string) (*Client, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}
