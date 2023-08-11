package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
)

type SetCourseRequest struct {
	Name         string         `json:"name"`
	Date         string         `json:"date" pattern:"YYYY-MM-DD HH:mm:ss"`
	Review       string         `json:"review"`
	Places       []string       `json:"places"`
	PlaceReviews []PlaceReview  `json:"place_reviews"`
	PlacePhotos  [][]PlacePhoto `json:"place_photos"`
	IsPublic     bool           `json:"is_public"`
}

type SetCourseResponse struct {
	Id string `json:"id"`
}

func AddCourse(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		claimsUserId, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(w, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(claimsUserId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var req SetCourseRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := courses.Add(container, userId, req.Name, req.Date, req.Review)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := courses.SetPlaces(container, id, req.Places); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i, v := range req.PlaceReviews {
			reviewId, err := place.CreateReview(container, id, req.Places[i], userId, v.Latitude, v.Longitude, v.Review)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			origins := make([]string, len(req.PlacePhotos[i]))
			thumbnails := make([]string, len(req.PlacePhotos[i]))
			for j, v := range req.PlacePhotos[i] {
				origins[j] = v.OriginUrl
				thumbnails[j] = v.ThumbnailUrl
			}
			if err := place.SetReviewPictures(container, reviewId, origins, thumbnails); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if err := courses.UpdateIsPublic(container, id, req.IsPublic); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type UpdateCourseRequest struct {
	CourseId string `json:"course_id"`
	SetCourseRequest
}

func UpdateCourse(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		claimsUserId, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(w, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(claimsUserId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var req UpdateCourseRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		courseId := sqlserver.UUID{}
		if err := courseId.From(req.CourseId); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := courses.UpdateName(container, courseId, req.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := courses.UpdateDate(container, courseId, req.Date); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := courses.UpdateReview(container, courseId, req.Review); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := courses.UpdateIsPublic(container, courseId, req.IsPublic); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := courses.SetPlaces(container, courseId, req.Places); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		prevReviewIds, err := place.GetReviewIds(container, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		step := 0
		for step = 0; step < len(prevReviewIds) && step < len(req.PlaceReviews); step++ {
			if err := place.UpdateReview(container, prevReviewIds[step], req.PlaceReviews[step].Review); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			origins := make([]string, len(req.PlacePhotos[step]))
			thumbnails := make([]string, len(req.PlacePhotos[step]))

			for j, v := range req.PlacePhotos[step] {
				origins[j] = v.OriginUrl
				thumbnails[j] = v.ThumbnailUrl
			}

			if err := place.SetReviewPictures(container, prevReviewIds[step], origins, thumbnails); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		for i := step; i < len(prevReviewIds); i++ {
			if err := place.DeleteReview(container, prevReviewIds[i]); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := place.DeleteReviewPictures(container, prevReviewIds[i]); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		for i := step; i < len(req.PlaceReviews); i++ {
			reviewId, err := place.CreateReview(container, courseId, req.Places[i], userId, req.PlaceReviews[i].Latitude, req.PlaceReviews[i].Longitude, req.PlaceReviews[i].Review)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			origins := make([]string, len(req.PlacePhotos[i]))
			thumbnails := make([]string, len(req.PlacePhotos[i]))
			for j, v := range req.PlacePhotos[i] {
				origins[j] = v.OriginUrl
				thumbnails[j] = v.ThumbnailUrl
			}

			if err := place.SetReviewPictures(container, reviewId, origins, thumbnails); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
