package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
)

func CreatePlaceReview(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		req := place.CreateReviewRequest{}

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

		authorUUID, err := base64.URLEncoding.DecodeString(req.AuthorId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		insertedId, err := place.CreateReview(container, courseUUID, req.PlaceId, authorUUID, req.Latitude, req.Longitude, req.Review)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := place.CreateReviewResponse{
			Id: base64.URLEncoding.EncodeToString(insertedId),
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetPlaceReviewsOfCourse(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		courseId := r.URL.Query().Get("course_id")

		courseUUID, err := base64.URLEncoding.DecodeString(courseId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		reviews, err := place.GetReviewsOfCourse(container, courseUUID)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := make([]place.GetReviewResponse, len(reviews))
		for i, review := range reviews {
			resp[i] = place.GetReviewResponse{
				Id:        base64.URLEncoding.EncodeToString(review.Id),
				AuthorId:  base64.URLEncoding.EncodeToString(review.AuthorId),
				Latitude:  review.Latitude,
				Longitude: review.Longitude,
				Review:    review.Review,
			}
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdatePlaceReview(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		req := place.UpdateReviewRequest{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := base64.URLEncoding.DecodeString(req.Id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := place.UpdateReview(container, id, req.Review); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeletePlaceReview(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		reviewUUID, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := place.DeleteReview(container, reviewUUID); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
