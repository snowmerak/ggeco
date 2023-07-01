package search

import (
	"encoding/json"
	"github.com/snowmerak/ggeco/gen/bean"
	"github.com/snowmerak/ggeco/lib/maps"
	"net/http"
	"strconv"
)

func Handler(container *bean.Container) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		radius, _ := strconv.ParseInt(r.URL.Query().Get("radius"), 10, 64)
		lang := r.URL.Query().Get("lang")
		latitude, _ := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
		longitude, _ := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
		openNow, _ := strconv.ParseBool(r.URL.Query().Get("opennow"))

		resp, err := maps.SearchText(r.Context(), container, func(request *maps.SearchTextRequest) *maps.SearchTextRequest {
			request.Query = query
			request.Radius = radius
			request.Language = lang
			request.Latitude = latitude
			request.Longitude = longitude
			request.OpenNow = openNow
			return request
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
