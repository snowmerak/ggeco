package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"net/http"
)

type GetBadgesRequest struct {
}

type Badge struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
	Image   string `json:"image,omitempty"`
}

type GetBadgesResponse struct {
	Badges []Badge `json:"badges"`
}

func GetBadges(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		list, err := badges.GetList(container)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		result := GetBadgesResponse{}
		for _, b := range list {
			result.Badges = append(result.Badges, Badge{
				Id:      b.Id.String(),
				Name:    b.Name,
				Summary: b.Summary,
				Image:   b.Image,
			})
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(result); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		wr.Header().Set("Content-Type", "application/json")
	}
}
