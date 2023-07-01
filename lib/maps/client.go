package maps

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"googlemaps.github.io/maps"
)

// jetti:bean Client
type Client struct {
	baseClient *maps.Client
	apiKey     string
	signature  []byte
}

func New(apiKey, signature string) (*Client, error) {
	sig, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return nil, err
	}
	c, err := maps.NewClient(maps.WithAPIKeyAndSignature(apiKey, signature))
	if err != nil {
		return nil, err
	}
	return &Client{c, apiKey, sig}, nil
}

func (c *Client) SignPhotoURL(photoRef string) string {
	mac := hmac.New(sha1.New, c.signature)
	path := "/maps/api/place/photo?maxwidth=400&photoreference=" + photoRef
	mac.Write([]byte(path))
	signature := mac.Sum(nil)
	return "https://maps.googleapis.com" + path + "&key=" + c.apiKey + "&signature=" + base64.URLEncoding.EncodeToString(signature)
}
