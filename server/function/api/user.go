package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/users"
	"net/http"
)

func GetUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := users.GetUser(container, uuid)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := users.GetUserResponse{
			Id:         base64.URLEncoding.EncodeToString(data.Id),
			Nickname:   data.Nickname,
			CreateAt:   data.CreateAt,
			LastSignin: data.LastSignin,
		}

		if data.Badge != nil {
			resp.Badge = base64.URLEncoding.EncodeToString(data.Badge)
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func AddUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req users.AddUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := users.AddUser(container, req.Nickname)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := users.AddUserResponse{
			Id: base64.URLEncoding.EncodeToString(id),
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateUserNickname(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req users.UpdateNicknameRequest

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		uuid, err := base64.URLEncoding.DecodeString(req.Id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := users.UpdateNickname(container, uuid, req.Nickname); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateUserLastSigninDate(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req users.UpdateLastSigninRequest

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		uuid, err := base64.URLEncoding.DecodeString(req.Id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := users.UpdateLastSignin(container, uuid); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := users.DeleteUser(container, uuid); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
