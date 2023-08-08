package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/users"
	"net/http"
	"strconv"
)

func GetKakaoUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := users.GetKakaoUser(container, uuid)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := users.GetKakaoUserResponse{
			Id:      id,
			KakaoId: user.KakaoId,
			Info:    user.Info,
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func GetKakaoUserByKakaoId(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		kakaoId := r.URL.Query().Get("kakao_id")

		convertedId, err := strconv.ParseInt(kakaoId, 10, 64)

		user, err := users.GetKakaoUserByKakaoId(container, convertedId)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := users.GetKakaoUserByKakaoIdResponse{
			Id:      base64.URLEncoding.EncodeToString(user.UserId),
			KakaoId: user.KakaoId,
			Info:    user.Info,
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func AddKakaoUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req users.AddKakaoUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			return
		}

		err := users.AddKakaoUser(container, req.UserId, req.KakaoId, req.Info)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func UpdateKakaoUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req users.UpdateKakaoUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			return
		}

		err := users.UpdateKakaoUser(container, req.Id, req.KakaoId, req.Info)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func DeleteKakaoUser(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("id")

		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			return
		}

		err = users.DeleteKakaoUser(container, uuid)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
