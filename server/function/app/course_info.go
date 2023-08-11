package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
)

type GetCourseInfoRequest struct {
	CourseId string `query:"course_id"`
}

type PlaceReview struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Review    string  `json:"review"`
}

type PlacePhoto struct {
	OriginUrl    string `json:"origin_url"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

type GetCourseInfoResponse struct {
	Course       Course                       `json:"course"`
	Places       []maps.SearchPlaceIdResponse `json:"places"`
	PlaceReviews []PlaceReview                `json:"place_reviews"`
	PlacePhotos  [][]PlacePhoto               `json:"place_photos"`
}

func GetCourseInfo(container bean.Container) httprouter.Handle {
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

		courseId := sqlserver.UUID{}
		if err := courseId.From(r.URL.Query().Get("course_id")); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := GetCourseInfoResponse{}

		course, err := courses.Get(container, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Course = Course{
			Id:       course.Id.String(),
			AuthorId: course.AuthorID.String(),
			Name:     course.Name,
			RegDate:  course.RegDate,
			Review:   course.Review,
		}

		favoriteCount, err := courses.CountFavoriteCourse(container, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Course.Favorites = favoriteCount

		isFavorite, err := courses.CheckFavoriteCourse(container, userId, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Course.IsFavorite = isFavorite

		placeList, err := courses.GetPlaces(container, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Places = make([]maps.SearchPlaceIdResponse, 0, len(placeList))
		for _, p := range placeList {
			placeInfo, err := place.GetPlace(container, p.PlaceId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			resp.Places = append(resp.Places, placeInfo.Data)
		}

		resp.PlaceReviews = make([]PlaceReview, 0, len(placeList))
		reviews, err := place.GetReviewsOfCourse(container, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, review := range reviews {
			resp.PlaceReviews = append(resp.PlaceReviews, PlaceReview{
				Latitude:  review.Latitude,
				Longitude: review.Longitude,
				Review:    review.Review,
			})
		}

		resp.PlacePhotos = make([][]PlacePhoto, 0, len(placeList))
		for _, review := range reviews {
			photos, err := place.GetReviewPictures(container, review.Id)
			if err != nil && errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			respPhotos := make([]PlacePhoto, 0, len(photos))
			for _, photo := range photos {
				respPhotos = append(respPhotos, PlacePhoto{
					OriginUrl:    photo.PictureUrl,
					ThumbnailUrl: photo.ThumbnailUrl,
				})
			}

			resp.PlacePhotos = append(resp.PlacePhotos, respPhotos)
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
