package maps

import (
	"context"
	"github.com/snowmerak/ggeco/gen/bean"
	"googlemaps.github.io/maps"
)

// SearchTextRequest is the request type of SearchText
// jetti:parameter
type SearchTextRequest struct {
	Query     string  `json:"query"`
	Radius    int64   `json:"radius"`
	Language  string  `json:"language"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	OpenNow   bool    `json:"open_now"`
}

// SearchTextResponse is the response type of SearchText
type SearchTextResponse struct {
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
	Name             string   `json:"name"`
	Icon             string   `json:"icon"`
	PlaceID          string   `json:"place_id"`
	Types            []string `json:"types"`
	BusinessStatus   string   `json:"business_status"`
	Rating           float64  `json:"rating,omitempty"`
	UserRatingsTotal int      `json:"user_ratings_total,omitempty"`
	OpeningHours     struct {
		OpenNow bool `json:"open_now"`
	} `json:"opening_hours,omitempty"`
	Photos     []PlacePhoto `json:"photos,omitempty"`
	PriceLevel int          `json:"price_level,omitempty"`
}

func SearchText(ctx context.Context, container *bean.Container, fn ...SearchTextRequestOptional) (response []*SearchTextResponse, err error) {
	client, err := GetClient(container)
	if err != nil {
		return response, err
	}

	request := ApplySearchTextRequest(SearchTextRequest{}, fn...)

	args := &maps.TextSearchRequest{
		Query:    request.Query,
		Language: "ko",
	}
	if request.Radius != 0 {
		args.Radius = uint(request.Radius)
	}
	if request.Language != "" {
		args.Language = request.Language
	}
	if request.Latitude != 0 && request.Longitude != 0 {
		args.Location = &maps.LatLng{
			Lat: request.Latitude,
			Lng: request.Longitude,
		}
	}

	resp, err := client.baseClient.TextSearch(ctx, args)
	if err != nil {
		return response, err
	}

	response = make([]*SearchTextResponse, len(resp.Results))
	for i, r := range resp.Results {
		response[i] = &SearchTextResponse{
			FormattedAddress: r.FormattedAddress,
			Geometry: struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
			}{
				Location: struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				}{
					Lat: r.Geometry.Location.Lat,
					Lng: r.Geometry.Location.Lng,
				},
			},
			Name:           r.Name,
			Icon:           r.Icon,
			PlaceID:        r.PlaceID,
			Types:          r.Types,
			BusinessStatus: r.BusinessStatus,
			Photos:         make([]PlacePhoto, len(r.Photos)),
			PriceLevel:     r.PriceLevel,
		}

		if r.OpeningHours != nil && r.OpeningHours.OpenNow != nil {
			response[i].OpeningHours.OpenNow = *r.OpeningHours.OpenNow
		}
		if r.Rating != 0 {
			response[i].Rating = float64(r.Rating)
		}
		if r.UserRatingsTotal != 0 {
			response[i].UserRatingsTotal = r.UserRatingsTotal
		}
		for j, p := range r.Photos {
			response[i].Photos[j] = PlacePhoto{
				Height:           p.Height,
				Width:            p.Width,
				HTMLAttributions: p.HTMLAttributions,
				PhotoReference:   p.PhotoReference,
			}
		}
	}

	return response, nil
}
