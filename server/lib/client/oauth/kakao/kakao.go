package kakao

import (
	"encoding/json"
	"net/http"
)

const uri = "https://kapi.kakao.com/v1/user/access_token_info"

type UserInfo struct {
	Id        int64 `json:"id"`
	ExpiresIn int   `json:"expires_in"`
	AppId     int   `json:"app_id"`
}

type Client struct {
	client http.Client
}

var client = &Client{
	client: http.Client{},
}

func NewClient() *Client {
	return client
}

func (c *Client) GetUserInfo(token string) (ui UserInfo, err error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return ui, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return ui, err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&ui)
	if err != nil {
		return ui, err
	}

	return ui, nil
}
