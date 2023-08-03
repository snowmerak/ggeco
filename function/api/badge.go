package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/gen/bean"
	"github.com/snowmerak/ggeco/lib/service/badges"
	"net/http"
)

func GetBadge(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		data, err := badges.Get(container, id)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(data); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		return
	}
}

func GetBadgeByName(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		name := r.URL.Query().Get("name")

		data, err := badges.GetByName(container, name)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(data); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		return
	}
}

func GetBadges(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data, err := badges.GetList(container)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(data); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		return
	}
}

func AddBadge(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req badges.AddBadgeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		if _, err := badges.Add(container, req.Name, req.Summary, req.Image); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		wr.WriteHeader(http.StatusOK)
	}
}

func UpdateBadgeName(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req badges.UpdateBadgeNameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		id, err := base64.URLEncoding.DecodeString(req.BadgeID)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		if err := badges.UpdateName(container, id, req.Name); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}
	}
}

func UpdateBadgeSummary(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req badges.UpdateBadgeSummaryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		id, err := base64.URLEncoding.DecodeString(req.BadgeID)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		if err := badges.UpdateSummary(container, id, req.Summary); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}
	}
}

func UpdateBadgeImage(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req badges.UpdateBadgeImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		id, err := base64.URLEncoding.DecodeString(req.BadgeID)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		if err := badges.UpdateImage(container, id, req.Image); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}
	}
}

func DeleteBadge(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			wr.Write([]byte(err.Error()))
			return
		}

		if err := badges.Delete(container, uuid); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			wr.Write([]byte(err.Error()))
			return
		}

		wr.WriteHeader(http.StatusOK)
	}
}
