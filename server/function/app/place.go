package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"net/http"
	"strconv"
)

type SearchPlacesRequest struct {
	Query  string  `query:"query"`
	Lat    float64 `query:"lat"`
	Lng    float64 `query:"lng"`
	Lang   string  `query:"lang"`
	Radius int     `query:"radius"`
}

type SearchPlacesResponse struct {
	Places []*maps.SearchTextResponse `json:"places"`
}

func SearchPlaces(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := r.URL.Query().Get("query")

		list, err := maps.SearchText(r.Context(), container, func(request *maps.SearchTextRequest) *maps.SearchTextRequest {
			request.Query = query
			lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
			if err == nil {
				request.Latitude = lat
			}
			lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
			if err == nil {
				request.Longitude = lng
			}
			request.Language = r.URL.Query().Get("lang")
			radius, err := strconv.Atoi(r.URL.Query().Get("radius"))
			if err == nil {
				request.Radius = int64(radius)
			}

			return request
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := SearchPlacesResponse{
			Places: list,
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
