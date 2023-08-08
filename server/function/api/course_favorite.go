package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"net/http"
)

func GetFavoriteCoursesByUserId(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("user_id")

		userId, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		favoriteCourses, err := courses.GetFavoriteCoursesByUserId(container, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var response []courses.GetFavoriteCourseByUserIdResponse
		for _, favoriteCourse := range favoriteCourses {
			response = append(response, courses.GetFavoriteCourseByUserIdResponse{
				Id:           base64.URLEncoding.EncodeToString(favoriteCourse.Id),
				UserId:       base64.URLEncoding.EncodeToString(favoriteCourse.UserId),
				CourseId:     base64.URLEncoding.EncodeToString(favoriteCourse.CourseId),
				RegisteredAt: favoriteCourse.RegisteredAt,
			})
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CountFavoriteCourse(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := r.URL.Query().Get("course_id")

		courseId, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		count, err := courses.CountFavoriteCourse(container, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := courses.CountFavoriteCourseResponse{
			Count: count,
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func AddFavoriteCourse(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req courses.AddFavoriteCourseRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userId, err := base64.URLEncoding.DecodeString(req.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		courseId, err := base64.URLEncoding.DecodeString(req.CourseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		insertedId, err := courses.AddFavoriteCourse(container, userId, courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := courses.AddFavoriteCourseResponse{
			Id: base64.URLEncoding.EncodeToString(insertedId),
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteFavoriteCourse(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := r.URL.Query().Get("id")

		favoriteCourseId, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = courses.DeleteFavoriteCourse(container, favoriteCourseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
