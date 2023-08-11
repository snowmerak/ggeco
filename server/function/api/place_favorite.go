package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
)

func GetFavoritePlacesByUserId(container sqlserver.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("user_id")

		userId, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		favoritePlaces, err := place.GetFavoritePlacesByUserId(container, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var response []place.GetFavoritePlaceByUserIdResponse
		for _, favoritePlace := range favoritePlaces {
			response = append(response, place.GetFavoritePlaceByUserIdResponse{
				Id:           base64.URLEncoding.EncodeToString(favoritePlace.Id),
				UserId:       base64.URLEncoding.EncodeToString(favoritePlace.UserId),
				PlaceId:      favoritePlace.PlaceId,
				RegisteredAt: favoritePlace.RegisteredAt,
			})
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CountFavoritePlace(container sqlserver.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := r.URL.Query().Get("place_id")

		count, err := place.CountFavoritePlace(container, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(count); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func AddFavoritePlace(container sqlserver.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req place.AddFavoritePlaceRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userId, err := base64.URLEncoding.DecodeString(req.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		favoritePlaceId, err := place.AddFavoritePlace(container, userId, req.PlaceId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := place.AddFavoritePlaceResponse{
			Id: base64.URLEncoding.EncodeToString(favoritePlaceId),
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

//func DeleteFavoritePlace(container sqlserver.Container) httprouter.Handle {
//	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//		id := r.URL.Query().Get("id")
//
//		favoritePlaceId, err := base64.URLEncoding.DecodeString(id)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//
//		if err := place.DeleteFavoritePlace(container, favoritePlaceId); err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
//}
