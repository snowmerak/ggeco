package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/lib/service/place"
	"net/http"
	"strconv"

	"github.com/snowmerak/ggeco/gen/bean"
)

func Search(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		query := r.URL.Query().Get("query")
		radius, _ := strconv.ParseInt(r.URL.Query().Get("radius"), 10, 64)
		lang := r.URL.Query().Get("lang")
		latitude, _ := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
		longitude, _ := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
		openNow, _ := strconv.ParseBool(r.URL.Query().Get("opennow"))

		resp, err := place.SearchText(r.Context(), container, query, radius, lang, latitude, longitude, openNow)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
