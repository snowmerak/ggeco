package naver

import (
	"encoding/json"
	"net/http"
)

const uri = "https://openapi.naver.com/v1/nid/me"

type UserInfo struct {
	ResultCode string `json:"resultcode"`
	Message    string `json:"message"`
	Response   struct {
		Id           string `json:"id"`
		Nickname     string `json:"nickname"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		Gender       string `json:"gender"`
		Age          string `json:"age"`
		Birthday     string `json:"birthday"`
		ProfileImage string `json:"profile_image"`
		BirthYear    string `json:"birthyear"`
		Mobile       string `json:"mobile"`
	} `json:"response"`
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
	req := http.Request{}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(&req)
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
