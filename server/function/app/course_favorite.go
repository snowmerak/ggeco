package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"net/http"
)

type IsFavoriteCourseRequest struct {
	CourseId string `query:"course_id"`
}

type IsFavoriteCourseResponse struct {
	IsFavorite bool `json:"is_favorite"`
}

func IsFavoriteCourse(container bean.Container) httprouter.Handle {
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
		courseIdStr := r.URL.Query().Get("course_id")
		if err := courseId.From(courseIdStr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		isFavorite, err := courses.CheckFavoriteCourse(container, userId, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := IsFavoriteCourseResponse{
			IsFavorite: isFavorite,
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type AddFavoriteCourseRequest struct {
	CourseId string `json:"course_id"`
}

func AddFavoriteCourse(container bean.Container) httprouter.Handle {
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

		var req AddFavoriteCourseRequest
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

		if _, err := courses.AddFavoriteCourse(container, userId, courseId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

type RemoveFavoriteCourseRequest struct {
	CourseId string `query:"course_id"`
}

func RemoveFavoriteCourse(container bean.Container) httprouter.Handle {
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

		if err := courses.DeleteFavoriteCourse(container, userId, courseId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type GetFavoriteCoursesResponse struct {
	Courses []Course `json:"courses"`
}

func GetFavoriteCourses(container bean.Container) httprouter.Handle {
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

		list, err := courses.GetFavoriteCoursesByUserId(container, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := GetFavoriteCoursesResponse{
			Courses: make([]Course, len(list)),
		}

		for i, v := range list {
			courseInfo, err := courses.Get(container, v.CourseId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			favoriteCount, err := courses.CountFavoriteCourse(container, v.CourseId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result.Courses[i] = Course{
				Id:         courseInfo.Id.String(),
				AuthorId:   courseInfo.AuthorID.String(),
				Name:       courseInfo.Name,
				RegDate:    courseInfo.RegDate,
				Review:     courseInfo.Review,
				Favorites:  favoriteCount,
				IsFavorite: true,
			}
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
