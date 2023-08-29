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
	"github.com/snowmerak/ggeco/server/lib/service/users"
	"net/http"
)

type GetProfileResponse struct {
	Id                  string `json:"id"`
	Nickname            string `json:"nickname"`
	BadgeId             string `json:"badge_id"`
	BadgeImage          string `json:"badge_image"`
	BadgeSummary        string `json:"badge_summary"`
	FavoritePlaceCount  int    `json:"favorite_place_count"`
	FavoriteCourseCount int    `json:"favorite_course_count"`
}

func GetProfile(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		claims, err := auth.GetJwtClaims(req.Context())
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		userIdStr, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(wr, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(userIdStr); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		userInfo, err := users.GetUser(container, userId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		badge, err := badges.Get(container, userInfo.Badge)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		favoritePlaceCount, err := place.CountFavoritePlaceByUserId(container, userId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		favoriteCourseCount, err := courses.CountFavoriteCourseByUserId(container, userId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := GetProfileResponse{
			Id:                  userInfo.Id.String(),
			Nickname:            userInfo.Nickname,
			BadgeId:             badge.Id.String(),
			BadgeImage:          badge.ActiveImage,
			BadgeSummary:        badge.Summary,
			FavoritePlaceCount:  favoritePlaceCount,
			FavoriteCourseCount: favoriteCourseCount,
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		wr.Header().Set("Content-Type", "application/json")
	}
}

type UpdateNicknameRequest struct {
	Nickname string `json:"nickname" required:"true"`
}

func UpdateNickname(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req UpdateNicknameRequest

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		userIdStr, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(wr, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(userIdStr); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := users.UpdateNickname(container, userId, req.Nickname); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type UpdateBadgeRequest struct {
	BadgeId string `json:"badge_id" required:"true"`
}

func UpdateBadge(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req UpdateBadgeRequest

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		userIdStr, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(wr, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(userIdStr); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		badgeId := sqlserver.UUID{}
		badgeIdStr, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(wr, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := badgeId.From(badgeIdStr); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := users.UpdateBadge(container, userId, badgeId); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
