package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
	"strconv"
)

type GetPlaceInfoRequest struct {
	PlaceId     string `query:"place_id"`
	CourseCount int    `query:"course_count"`
}

type GetPlaceInfoResponse struct {
	Data          maps.SearchPlaceIdResponse `json:"data"`
	Courses       []Course                   `json:"courses"`
	FavoriteCount int                        `json:"favorite_count"`
}

func GetPlaceInfo(container bean.Container) httprouter.Handle {
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

		placeId := r.URL.Query().Get("place_id")

		courseCount, err := strconv.Atoi(r.URL.Query().Get("course_count"))
		if err != nil {
			courseCount = 3
		}

		resp := GetPlaceInfoResponse{}
		resp.Courses = make([]Course, 0)

		placeInfo, err := place.GetPlace(container, placeId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Data = placeInfo.Data

		favoriteCount, err := place.CountFavoritePlace(container, placeId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.FavoriteCount = favoriteCount

		isFavorite, err := place.CheckFavoritePlace(container, userId, placeId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Data.IsFavorite = isFavorite

		courseIds, err := place.GetCourseContainsPlace(container, placeId, courseCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, courseId := range courseIds {
			course, err := courses.Get(container, courseId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			resp.Courses = append(resp.Courses, Course{
				Id:       course.Id.String(),
				AuthorId: course.AuthorID.String(),
				Name:     course.Name,
				RegDate:  course.RegDate,
				Review:   course.Review,
			})
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}
