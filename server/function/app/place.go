package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
	"strconv"
)

type SearchPlacesRequest struct {
	Query  string  `query:"query"`
	Lat    float64 `query:"lat"`
	Lng    float64 `query:"lng"`
	Lang   string  `query:"lang"`
	Radius int     `query:"radius"`
}

type SearchPlacesResponse struct {
	Places []*maps.SearchTextResponse `json:"places"`
}

func SearchPlaces(container bean.Container) httprouter.Handle {
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

		query := r.URL.Query().Get("query")

		list, err := maps.SearchText(r.Context(), container, func(request *maps.SearchTextRequest) *maps.SearchTextRequest {
			request.Query = query
			lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
			if err == nil {
				request.Latitude = lat
			}
			lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
			if err == nil {
				request.Longitude = lng
			}
			request.Language = r.URL.Query().Get("lang")
			radius, err := strconv.Atoi(r.URL.Query().Get("radius"))
			if err == nil {
				request.Radius = int64(radius)
			}

			return request
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := SearchPlacesResponse{
			Places: list,
		}

		for _, p := range result.Places {
			isFavorite, err := place.CheckFavoritePlace(container, userId, p.PlaceID)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			p.IsFavorite = isFavorite
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
