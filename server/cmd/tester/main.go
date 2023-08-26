package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/snowmerak/ggeco/server/cmd/tester/internal/client"
	"github.com/snowmerak/ggeco/server/function/app"
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
	at := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkX2F0IjoxNjkzNjM1NDE2LCJraW5kIjowLCJ1c2VyX2lkIjoiS0NiZEJrUFlFZTY2RndEX0VVd1gxUT09IiwidXNlcl9uaWNrIjoidGVzdCJ9.iX5xQms2-lD2E8Bzo-3NPvRnjCQBcyrxsRFNzcZl1FDpU6YoOncY1wnbWcwdqEJSZreTXUBiV5KpFsZsb7VAPQ"
	resp, err := cli.GetPlaceInfo(at, app.GetPlaceInfoRequest{
		PlaceId: "ChIJJ0hjaMKYfDUR3vuPnM6MgS8",
	})
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(resp); err != nil {
		panic(err)
	}
}
