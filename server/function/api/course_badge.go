package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"github.com/snowmerak/ggeco/server/lib/service/courses"
	"net/http"
)

func GetCourseBadges(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("course_id")
		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := courses.GetCourseBadges(container, uuid)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := make([]courses.GetCourseBadgeResponse, len(data))
		for i, v := range data {
			badge, err := badges.Get(container, v.BadgeId)
			if err != nil {
				http.Error(wr, err.Error(), http.StatusInternalServerError)
				return
			}

			resp[i].BadgeId = base64.URLEncoding.EncodeToString(badge.Id)
			resp[i].BadgeName = badge.Name
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func SetCourseBadges(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		req := courses.SetCourseBadgesRequest{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		courseUUID, err := base64.URLEncoding.DecodeString(req.CourseId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		badgeUUIDs := make([]sqlserver.UUID, len(req.BadgeIds))
		for i, v := range req.BadgeIds {
			badgeUUIDs[i], err = base64.URLEncoding.DecodeString(v)
			if err != nil {
				http.Error(wr, err.Error(), http.StatusBadRequest)
				return
			}
		}

		if err := courses.SetCourseBadges(container, courseUUID, badgeUUIDs); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
