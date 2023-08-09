package app

import (
	"context"
	"errors"
	"github.com/snowmerak/ggeco/server/function/api"
)

var ErrApiKeyNotFound = errors.New("api key not found")
var ErrAddressNotFound = errors.New("address not found")

func GetApiInfo(ctx context.Context) (apiKey string, address string, err error) {
	apiKey, ok := ctx.Value(api.AzureApiKey{}).(string)
	if !ok {
		return "", "", ErrApiKeyNotFound
	}
	address, ok = ctx.Value(api.Address{}).(string)
	if !ok {
		return "", "", ErrAddressNotFound
	}
	return apiKey, address, nil
}
