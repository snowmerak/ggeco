package place

import (
	"context"
	"github.com/snowmerak/ggeco/lib/client/maps"
)

func SearchText(ctx context.Context, container maps.Container, query string, radius int64, lang string, latitude float64, longitude float64, openNow bool) (response []*maps.SearchTextResponse, err error) {
	resp, err := maps.SearchText(ctx, container, func(request *maps.SearchTextRequest) *maps.SearchTextRequest {
		request.Query = query
		request.Radius = radius
		request.Language = lang
		request.Latitude = latitude
		request.Longitude = longitude
		request.OpenNow = openNow
		return request
	})
	if err != nil {
		return response, err
	}

	return resp, nil
}
