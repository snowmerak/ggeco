package client

import (
	"encoding/json"
	"github.com/snowmerak/ggeco/server/function/app"
)

func (c *Client) SearchPlaces(at string, req app.SearchPlacesRequest) (app.SearchPlacesResponse, error) {
	data, err := c.Get("/place/search", at, c.Query("query", req.Query))
	if err != nil {
		return app.SearchPlacesResponse{}, err
	}

	resp := app.SearchPlacesResponse{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) GetPlaceInfo(at string, req app.GetPlaceInfoRequest) (app.GetPlaceInfoResponse, error) {
	data, err := c.Get("/place", at, c.Query("place_id", req.PlaceId))
	if err != nil {
		return app.GetPlaceInfoResponse{}, err
	}

	resp := app.GetPlaceInfoResponse{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
