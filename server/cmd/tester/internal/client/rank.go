package client

import (
	"encoding/json"
	"github.com/snowmerak/ggeco/server/function/app"
)

func (c *Client) GetMyBadgeRank(at string) (app.GetBadgeRankResponse, error) {
	data, err := c.Get("/badge/rank", at)
	if err != nil {
		return app.GetBadgeRankResponse{}, err
	}

	resp := app.GetBadgeRankResponse{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
