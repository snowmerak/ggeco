package app

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/oauth/kakao"
	"github.com/snowmerak/ggeco/server/lib/client/oauth/naver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/names"
	"github.com/snowmerak/ggeco/server/lib/service/users"
	"net/http"
)

type SignInRequest struct {
	NaverAccount bool   `json:"naver_account,omitempty"`
	KakaoAccount bool   `json:"kakao_account,omitempty"`
	AccessToken  string `json:"access_token" required:"true"`
}

type SignInResponse struct {
	RefreshToken string `json:"refresh_token"`
}

func SignIn(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		var reqBody SignInRequest
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&reqBody)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		var jwtToken string
		switch {
		case reqBody.NaverAccount:
			jwtToken, err = signInNaver(req.Context(), container, reqBody.AccessToken)
			if err != nil {
				log.Error().Err(err).Msg("sign in naver")
				http.Error(wr, err.Error(), http.StatusInternalServerError)
				return
			}
		case reqBody.KakaoAccount:
			jwtToken, err = signInKakao(req.Context(), container, reqBody.AccessToken)
			if err != nil {
				log.Error().Err(err).Msg("sign in kakao")
				http.Error(wr, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			http.Error(wr, "No Resource Owner Selection", http.StatusBadRequest)
			wr.WriteHeader(http.StatusBadRequest)
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(SignInResponse{RefreshToken: jwtToken}); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func signInNaver(ctx context.Context, container bean.Container, token string) (string, error) {
	jwtSecretKey, err := auth.GetJwtSecretKey(ctx)
	if err != nil {
		return "", err
	}

	client := naver.NewClient()

	ui, err := client.GetUserInfo(token)
	if err != nil {
		return "", err
	}

	rs, err := users.GetNaverUserByNaverId(container, ui.Response.Id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return "", err
		}

		userId, err := users.AddUser(container, names.MakeNewName())
		if err != nil {
			return "", err
		}

		buf, err := json.Marshal(ui.Response)
		if err != nil {
			return "", err
		}

		info := base64.URLEncoding.EncodeToString(buf)

		if err := users.AddNaverUser(container, userId, ui.Response.Id, info); err != nil {
			return "", err
		}

		rs.NaverId = ui.Response.Id
		rs.Id = userId
		rs.Info = info
		log.Info().Msgf("new user: %d", userId)
	}

	user, err := users.GetUser(container, rs.Id)
	if err != nil {
		return "", err
	}

	jwtToken, err := auth.MakeUserToken(jwtSecretKey, base64.URLEncoding.EncodeToString(user.Id), user.Nickname, auth.RefreshTokenLifetime())
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func signInKakao(ctx context.Context, container bean.Container, token string) (string, error) {
	jwtSecretKey, err := auth.GetJwtSecretKey(ctx)
	if err != nil {
		return "", err
	}

	client := kakao.NewClient()

	ui, err := client.GetUserInfo(token)
	if err != nil {
		return "", err
	}

	rs, err := users.GetKakaoUserByKakaoId(container, ui.Id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return "", err
		}

		userId, err := users.AddUser(container, names.MakeNewName())
		if err != nil {
			return "", err
		}

		buf, err := json.Marshal(ui)
		if err != nil {
			return "", err
		}

		info := base64.URLEncoding.EncodeToString(buf)

		if err := users.AddKakaoUser(container, userId, ui.Id, info); err != nil {
			return "", err
		}

		rs.KakaoId = ui.Id
		rs.UserId = userId
		rs.Info = info
	}

	user, err := users.GetUser(container, rs.UserId)
	if err != nil {
		return "", err
	}

	jwtToken, err := auth.MakeUserToken(jwtSecretKey, base64.URLEncoding.EncodeToString(user.Id), user.Nickname, auth.RefreshTokenLifetime())
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
