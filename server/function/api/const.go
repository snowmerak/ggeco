package api

import "context"

type AzureApiKey struct{}

func GetApiKey(ctx context.Context) string {
	value, ok := ctx.Value(AzureApiKey{}).(string)
	if !ok {
		return ""
	}
	return value
}

func SetApiKey(ctx context.Context, apiKey string) context.Context {
	return context.WithValue(ctx, AzureApiKey{}, apiKey)
}

type Address struct{}

func GetAddress(ctx context.Context) string {
	value, ok := ctx.Value(Address{}).(string)
	if !ok {
		return ""
	}
	return value
}

func SetAddress(ctx context.Context, address string) context.Context {
	return context.WithValue(ctx, Address{}, address)
}
