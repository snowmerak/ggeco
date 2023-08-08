package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/users"
	"net/http"
)

func GetNaverUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := users.GetNaverUser(container, uuid)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := users.GetNaverUserResponse{
			Id:      base64.URLEncoding.EncodeToString(data.Id),
			NaverId: data.NaverId,
			Info:    data.Info,
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetNaverUserByNaverId(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		naverId := r.URL.Query().Get("naver_id")

		data, err := users.GetNaverUserByNaverId(container, naverId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := users.GetNaverUserByNaverIdResponse{
			Id:      base64.URLEncoding.EncodeToString(data.Id),
			NaverId: data.NaverId,
			Info:    data.Info,
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func AddNaverUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req users.AddNaverUserRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := base64.URLEncoding.DecodeString(req.Id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := users.AddNaverUser(container, id, req.NaverId, req.Info); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateNaverUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req users.UpdateNaverUserRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := base64.URLEncoding.DecodeString(req.Id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := users.UpdateNaverUser(container, id, req.NaverId, req.Info); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteNaverUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := users.DeleteNaverUser(container, uuid); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
