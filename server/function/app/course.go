package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"net/http"
	"strconv"
)

type GetPopularCourseOfBadgeRequest struct {
	BadgeId string `query:"badge_id"`
	Count   int    `query:"count"`
}

type Course struct {
	Id               string `json:"id"`
	AuthorId         string `json:"author_id"`
	AuthorNickname   string `json:"author_nickname,omitempty"`
	AuthorBadgeImage string `json:"author_badge_image,omitempty"`
	AuthorBadgeName  string `json:"author_badge_name,omitempty"`
	Name             string `json:"name"`
	RegDate          string `json:"reg_date"`
	Review           string `json:"review"`
	Favorites        int    `json:"favorites"`
	IsFavorite       bool   `json:"is_favorite"`
	VillageAddress   string `json:"village_address,omitempty"`
	Category         string `json:"category,omitempty"`
	TitleImage       string `json:"title_image,omitempty"`
}

type GetPopularCourseOfBadgeResponse struct {
	Courses []Course `json:"courses"`
}

func GetPopularCourseOfBadge(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		badgeId := new(sqlserver.UUID)
		if err := badgeId.From(r.URL.Query().Get("badge_id")); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		list, err := courses.GetPopularInBadge(container, *badgeId, count)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := GetPopularCourseOfBadgeResponse{}
		for _, v := range list {
			favoriteCount, err := courses.CountFavoriteCourse(container, v.Id)
			if err != nil && errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result.Courses = append(result.Courses, Course{
				Id:        v.Id.String(),
				AuthorId:  v.AuthorID.String(),
				Name:      v.Name,
				RegDate:   v.RegDate,
				Review:    v.Review,
				Favorites: favoriteCount,
			})
		}

		enc := json.NewEncoder(w)
		if err := enc.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

type GetRecentCoursesRequest struct {
	Count int `query:"count"`
}

type GetRecentCoursesResponse struct {
	Courses []Course `json:"courses"`
}

func GetRecentCourses(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		list, err := courses.GetNewest(container, count)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := GetRecentCoursesResponse{}
		for _, v := range list {
			favoriteCount, err := courses.CountFavoriteCourse(container, v.Id)
			if err != nil && errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result.Courses = append(result.Courses, Course{
				Id:        v.Id.String(),
				AuthorId:  v.AuthorID.String(),
				Name:      v.Name,
				RegDate:   v.RegDate,
				Review:    v.Review,
				Favorites: favoriteCount,
			})
		}

		enc := json.NewEncoder(w)
		if err := enc.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

type GetMyCoursesResponse struct {
	Courses []Course `json:"courses"`
}

func GetMyCourses(container bean.Container) httprouter.Handle {
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

		list, err := courses.GetByAuthor(container, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := GetMyCoursesResponse{}
		for _, v := range list {
			favoriteCount, err := courses.CountFavoriteCourse(container, v.Id)
			if err != nil && errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result.Courses = append(result.Courses, Course{
				Id:        v.Id.String(),
				AuthorId:  v.AuthorID.String(),
				Name:      v.Name,
				RegDate:   v.RegDate,
				Review:    v.Review,
				Favorites: favoriteCount,
			})
		}

		enc := json.NewEncoder(w)
		if err := enc.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

type GetRecommendedCoursesRequest struct {
	Count int `query:"count"`
}

type GetRecommendedCoursesResponse struct {
	Courses []Course `json:"courses"`
}

func GetRecommendedCourses(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			count = 10
		}

		list, err := courses.GetRandom(container, count)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := GetRecommendedCoursesResponse{}
		for _, v := range list {
			favoriteCount, err := courses.CountFavoriteCourse(container, v.Id)
			if err != nil && errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result.Courses = append(result.Courses, Course{
				Id:        v.Id.String(),
				AuthorId:  v.AuthorID.String(),
				Name:      v.Name,
				RegDate:   v.RegDate,
				Review:    v.Review,
				Favorites: favoriteCount,
			})
		}

		enc := json.NewEncoder(w)
		if err := enc.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}
