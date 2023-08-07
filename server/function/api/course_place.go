package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
)

func GetCoursePlaces(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		courseId := r.URL.Query().Get("course_id")

		courseUUID, err := base64.URLEncoding.DecodeString(courseId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := courses.GetPlaces(container, courseUUID)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := courses.GetPlacesResponse{
			Places: make([]place.GetPlaceResponse, len(data)),
		}
		for i, v := range data {
			resp.Places[i].ID = base64.URLEncoding.EncodeToString(v.Id)
			value, err := place.GetPlace(container, v.PlaceId)
			if err != nil {
				http.Error(wr, err.Error(), http.StatusInternalServerError)
				return
			}
			value.ID = v.PlaceId
			resp.Places[i] = value
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func SetCoursePlaces(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		req := courses.SetPlacesRequest{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		courseUUID, err := base64.URLEncoding.DecodeString(req.CourseId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		placeIds := make([]sqlserver.UUID, len(req.PlaceIds))
		for i, v := range req.PlaceIds {
			placeIds[i], err = base64.URLEncoding.DecodeString(v)
			if err != nil {
				http.Error(wr, err.Error(), http.StatusBadRequest)
				return
			}
		}

		if err := courses.SetPlaces(container, courseUUID, placeIds); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
