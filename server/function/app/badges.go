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
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Summary       string `json:"summary,omitempty"`
	ActiveImage   string `json:"active_image,omitempty"`
	InactiveImage string `json:"inactive_image,omitempty"`
	SelectedImage string `json:"selected_image,omitempty"`
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
				Id:            b.Id.String(),
				Name:          b.Name,
				Summary:       b.Summary,
				ActiveImage:   b.ActiveImage,
				InactiveImage: b.InactiveImage,
				SelectedImage: b.SelectedImage,
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

type GetSearchableBadgesRequest struct {
}

type GetSearchableBadgesResponse struct {
	Badges []Badge `json:"badges"`
}

func GetSearchableBadges(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		list, err := badges.GetSearchables(container)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		result := GetSearchableBadgesResponse{}
		for _, b := range list {
			result.Badges = append(result.Badges, Badge{
				Id:            b.Id.String(),
				Name:          b.Name,
				Summary:       b.Summary,
				ActiveImage:   b.ActiveImage,
				InactiveImage: b.InactiveImage,
				SelectedImage: b.SelectedImage,
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
