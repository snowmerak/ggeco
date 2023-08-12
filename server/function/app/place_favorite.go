package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
)

type IsFavoritePlaceRequest struct {
	PlaceId string `query:"place_id"`
}

type IsFavoritePlaceResponse struct {
	IsFavorite bool `json:"is_favorite"`
}

func IsFavoritePlace(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		claimsUserId, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(w, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(claimsUserId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		placeId := r.URL.Query().Get("place_id")

		isFavorite, err := place.CheckFavoritePlace(container, userId, placeId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := IsFavoritePlaceResponse{
			IsFavorite: isFavorite,
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type AddFavoritePlaceRequest struct {
	PlaceId string `json:"place_id"`
}

func AddFavoritePlace(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		claimsUserId, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(w, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(claimsUserId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var req AddFavoritePlaceRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := place.AddFavoritePlace(container, userId, req.PlaceId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

type RemoveFavoritePlaceRequest struct {
	PlaceId string `query:"place_id"`
}

func RemoveFavoritePlace(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		claimsUserId, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(w, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(claimsUserId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		placeId := r.URL.Query().Get("place_id")

		if err := place.DeleteFavoritePlace(container, userId, placeId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type GetFavoritePlacesResponse struct {
	Places []maps.SearchPlaceIdResponse `json:"places"`
}

func GetFavoritePlaces(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		claimsUserId, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(w, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(claimsUserId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		list, err := place.GetFavoritePlacesByUserId(container, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := GetFavoritePlacesResponse{
			Places: make([]maps.SearchPlaceIdResponse, len(list)),
		}

		for i, favoritePlace := range list {
			p, err := place.GetPlace(container, favoritePlace.PlaceId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			result.Places[i] = p.Data
			result.Places[i].IsFavorite = true
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
