package client

import (
	"encoding/json"
	"github.com/snowmerak/ggeco/server/function/app"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
)

func (c *Client) MakeNewUser(secret []byte, userId, userNick string) (rt, at string, err error) {
	rt, err = auth.MakeUserToken(secret, userId, userNick, auth.RefreshTokenLifetime())
	if err != nil {
		return "", "", err
	}

	at, err = auth.MakeUserToken(secret, userId, userNick, auth.AccessTokenLifetime())
	if err != nil {
		return "", "", err
	}

	return rt, at, nil
}

func (c *Client) RefreshToken(rt string) (string, error) {
	data, err := c.Post("/auth/refresh", rt, nil)
	if err != nil {
		return "", err
	}

	resp := app.RefreshResponse{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
