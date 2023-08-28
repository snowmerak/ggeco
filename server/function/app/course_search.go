package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"math/rand"
	"net/http"
	"strconv"
)

type FindCoursesBySearchPlaceRequest struct {
	Query  string  `query:"query" required:"true"`
	Lat    float64 `query:"lat"`
	Lng    float64 `query:"lng"`
	Lang   string  `query:"lang"`
	Radius int     `query:"radius"`
	Count  int     `query:"count"`
}

type FindCoursesBySearchPlaceResponse struct {
	Courses []Course `json:"courses"`
}

func FindCoursesBySearchPlace(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := r.URL.Query().Get("query")

		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			count = 10
		}

		places, err := maps.SearchText(r.Context(), container, func(request *maps.SearchTextRequest) *maps.SearchTextRequest {
			request.Query = query
			lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
			if err == nil {
				request.Latitude = lat
			}
			lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
			if err == nil {
				request.Longitude = lng
			}
			request.Language = r.URL.Query().Get("lang")
			radius, err := strconv.Atoi(r.URL.Query().Get("radius"))
			if err == nil {
				request.Radius = int64(radius)
			}

			return request
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(places)/2 >= count {
			places = places[:count/2]
		}

		result := FindCoursesBySearchPlaceResponse{}
		for _, place := range places {
			list, err := courses.GetCoursesFromPlace(container, place.PlaceID, count/len(places)+1)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

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
		}

		for i := range result.Courses {
			result.Courses[i], result.Courses[rand.Int()/len(places)] = result.Courses[rand.Int()/len(places)], result.Courses[i]
		}
		if len(result.Courses) > count {
			result.Courses = result.Courses[:count]
		}
		if len(result.Courses) == 0 {
			result.Courses = []Course{}
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}
