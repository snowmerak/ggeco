package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/snowmerak/ggeco/server/cmd/tester/internal/client"
	"github.com/unsafe-risk/dotenv"
	"os"
	"strings"
)

func main() {
	env, err := dotenv.Read(".test.env")
	if err != nil {
		panic(err)
	}

	url := strings.TrimSuffix(env["URL"], "\r")
	secret := strings.TrimSuffix(env["JWT_SECRET"], "\r")
	jwtSecretKey, err := base64.URLEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}
	_ = jwtSecretKey

	cli := client.New(url)

	// ChIJJ0hjaMKYfDUR3vuPnM6MgS8
	at := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkX2F0IjoxNjkzOTI4MTk4LCJraW5kIjowLCJ1c2VyX2lkIjoiNXdMT3k0ckFTUk9RdHUxRGU5TzlIQT09IiwidXNlcl9uaWNrIjoi6rO17IaQ7ZWcIOq5qOy9lCAxOTA5M-2YuCJ9.UY7nYk7h8DAl_NPT_07rz0hfuOfavb1MgBIO2eIW3WHAQIXYPo-BS5IdZbNHQHnS0KuXpphFHUw5FmySfAar6Q"
	resp, err := cli.GetMyBadgeRank(at)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./test.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(resp); err != nil {
		panic(err)
	}
}
