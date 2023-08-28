package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"net/http"
	"strings"
)

type RefreshRequest struct {
	RefreshToken string `header:"Authorization" required:"true"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func Refresh(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		jwtSecretKey, err := auth.GetJwtSecretKey(req.Context())
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") || len(authorizationHeader) < 8 {
			http.Error(wr, "No Authorization Header", http.StatusBadRequest)
			return
		}

		refreshToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

		token, err := auth.ValidateUserToken(jwtSecretKey, refreshToken)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		userId := token[auth.UserId]
		nickName := token[auth.UserNick]

		kind, ok := token[auth.Kind].(float64)
		if !ok {
			http.Error(wr, "Invalid Kind", http.StatusInternalServerError)
			return
		}

		if kind != float64(auth.KindRefreshToken) {
			http.Error(wr, "Invalid Kind", http.StatusBadRequest)
			return
		}

		accessToken, err := auth.MakeUserToken(jwtSecretKey, userId.(string), nickName.(string), auth.AccessTokenLifetime())
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(RefreshResponse{AccessToken: accessToken}); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		wr.Header().Set("Content-Type", "application/json")
	}
}
