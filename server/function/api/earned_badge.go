package api

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"net/http"
)

func AddEarnedBadge(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var req badges.AddEarnedBadgeRequest
		if err := decoder.Decode(&req); err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		userId, err := base64.URLEncoding.DecodeString(req.UserID)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		badgeId, err := base64.URLEncoding.DecodeString(req.BadgeID)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		if err := badges.AddEarnedBadge(container, userId, badgeId); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}
	}
}

func GetEarnedBadgesByUserId(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		userId := r.URL.Query().Get("user_id")

		uuid, err := base64.URLEncoding.DecodeString(userId)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		data, err := badges.GetEarnedBadgesByUserId(container, uuid)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				wr.WriteHeader(http.StatusNotFound)
				return
			}

			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		resp := make([]badges.GetEarnedBadgeResponse, 0, len(data))
		for i, v := range data {
			resp[i] = badges.GetEarnedBadgeResponse{
				Id:       base64.URLEncoding.EncodeToString(v.Id),
				BadgeID:  base64.URLEncoding.EncodeToString(v.BadgeId),
				UserID:   base64.URLEncoding.EncodeToString(v.UserId),
				EarnedAt: v.EarnedAt,
			}
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}
	}
}

func CountUsersEarnedBadge(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		badgeId := r.URL.Query().Get("badge_id")

		uuid, err := base64.URLEncoding.DecodeString(badgeId)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		data, err := badges.CountUsersEarnedBadge(container, uuid)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				wr.WriteHeader(http.StatusNotFound)
				return
			}

			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		resp := badges.CountUsersEarnedBadgeResponse{
			Count: data,
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}
	}
}
