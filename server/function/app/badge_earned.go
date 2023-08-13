package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"net/http"
)

type EarnedBadge struct {
	IsEarned   bool    `json:"is_earned"`
	BadgeId    string  `json:"badge_id"`
	Name       string  `json:"name"`
	Summary    string  `json:"summary"`
	Image      string  `json:"image"`
	EarnedRate float64 `json:"earned_rate"`
	EarnedAt   string  `json:"earned_at"`
}

type GetEarnedBadgesRequest struct {
}

type GetEarnedBadgesResponse struct {
	EarnedBadges []EarnedBadge `json:"earned_badges"`
}

func GetEarnedBadges(container bean.Container) httprouter.Handle {
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

		earnedBadges, err := badges.GetEarnedBadgesByUserId(container, userId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := GetEarnedBadgesResponse{}

		for _, earnedBadge := range earnedBadges {
			badge, err := badges.Get(container, earnedBadge.BadgeId)
			if err != nil {
				http.Error(wr, err.Error(), http.StatusInternalServerError)
				return
			}

			earnedRate, err := badges.GetEarnedRateOfBadge(container, earnedBadge.BadgeId)
			if err != nil {
				http.Error(wr, err.Error(), http.StatusInternalServerError)
				return
			}

			resp.EarnedBadges = append(resp.EarnedBadges, EarnedBadge{
				IsEarned:   true,
				BadgeId:    earnedBadge.Id.String(),
				Name:       badge.Name,
				Summary:    badge.Summary,
				Image:      badge.Image,
				EarnedRate: earnedRate,
				EarnedAt:   earnedBadge.EarnedAt,
			})
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type GetBadgeInfoRequest struct {
	BadgeId string `query:"badge_id"`
}

type GetBadgeInfoResponse struct {
	EarnedBadge EarnedBadge `json:"earned_badge"`
}

func GetBadgeInfo(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		badgeId := sqlserver.UUID{}
		badgeIdStr := req.URL.Query().Get("badge_id")
		if err := badgeId.From(badgeIdStr); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

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

		badge, err := badges.Get(container, badgeId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		isEarned, earnedAt, err := badges.CheckEarnedBadge(container, userId, badgeId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		earnedRate, err := badges.GetEarnedRateOfBadge(container, badgeId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := GetBadgeInfoResponse{
			EarnedBadge: EarnedBadge{
				IsEarned:   isEarned,
				BadgeId:    badge.Id.String(),
				Name:       badge.Name,
				Summary:    badge.Summary,
				Image:      badge.Image,
				EarnedRate: earnedRate,
				EarnedAt:   earnedAt,
			},
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
