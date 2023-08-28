package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
)

type SetCourseRequest struct {
	Name         string        `json:"name"`
	Date         string        `json:"date" pattern:"YYYY-MM-DD HH:mm:ss"`
	Review       string        `json:"review"`
	Places       []string      `json:"places"`
	PlaceReviews []PlaceReview `json:"place_reviews"`
	IsPublic     bool          `json:"is_public"`
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

		id, err := courses.Add(container, userId, req.Name, req.Date, req.Review, req.IsPublic)
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

			origins := make([]string, len(v.Photos))
			thumbnails := make([]string, len(v.Photos))
			for j, v := range v.Photos {
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

		{
			placeTypeSet := make(map[string]struct{})

			for i := range req.Places {
				p, err := place.GetPlace(container, req.Places[i])
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				for _, placeType := range p.Data.Types {
					if placeType == "" {
						continue
					}
					if err := place.AddOrInitVisitCount(container, userId, placeType); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					placeTypeSet[placeType] = struct{}{}
				}
			}

			badgeSet := make(map[string]struct{})
			badgeList := make([]sqlserver.UUID, 0, len(badgeSet))
			for placeType := range placeTypeSet {
				badge, err := badges.GetBadgeFromPlaceType(container, placeType)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				badgeName := badge.String()
				if badgeName == "" {
					continue
				}
				if _, ok := badgeSet[badge.String()]; !ok {
					badgeSet[badge.String()] = struct{}{}
					badgeList = append(badgeList, badge)
				}
			}

			if err := courses.SetCourseBadges(container, userId, badgeList); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		enc := json.NewEncoder(w)
		if err := enc.Encode(SetCourseResponse{Id: id.String()}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

type UpdateCourseRequest struct {
	CourseId string `query:"course_id"`
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

		courseId := sqlserver.UUID{}
		if err := courseId.From(r.URL.Query().Get("course_id")); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req UpdateCourseRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
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

			origins := make([]string, len(req.PlaceReviews[step].Photos))
			thumbnails := make([]string, len(req.PlaceReviews[step].Photos))

			for j, v := range req.PlaceReviews[step].Photos {
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

			origins := make([]string, len(req.PlaceReviews[i].Photos))
			thumbnails := make([]string, len(req.PlaceReviews[i].Photos))
			for j, v := range req.PlaceReviews[i].Photos {
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
