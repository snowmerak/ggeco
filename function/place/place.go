package place

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/snowmerak/ggeco/gen/bean"
	"github.com/snowmerak/ggeco/lib/maps"
)

func Handler(container *bean.Container) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		placeId := r.URL.Query().Get("place_id")
		lang := r.URL.Query().Get("lang")
		getReviews, _ := strconv.ParseBool(r.URL.Query().Get("reviews"))
		getOpeningHours, _ := strconv.ParseBool(r.URL.Query().Get("opening_hours"))
		getPermanentlyClosed, _ := strconv.ParseBool(r.URL.Query().Get("business_status"))
		getGeometryLocation, _ := strconv.ParseBool(r.URL.Query().Get("geometry_location"))
		getPhotos, _ := strconv.ParseBool(r.URL.Query().Get("photo"))
		getPhone, _ := strconv.ParseBool(r.URL.Query().Get("phone"))

		resp, err := maps.SearchPlaceId(r.Context(), container, func(request *maps.SearchPlaceIdRequest) *maps.SearchPlaceIdRequest {
			request.PlaceID = placeId
			request.Language = lang
			request.GetReviews = getReviews
			request.GetOpeningHours = getOpeningHours
			request.GetPermanentlyClosed = getPermanentlyClosed
			request.GetGeometryLocation = getGeometryLocation
			request.GetPhotos = getPhotos
			request.GetPhone = getPhone
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
