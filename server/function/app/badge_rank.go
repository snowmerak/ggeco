package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"net/http"
)

type GetBadgeRankResponse struct {
	Rank    int64  `json:"rank"`
	Delta   int64  `json:"delta"`
	Updated string `json:"updated"`
}

func GetBadgeRank(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		claimsUserId, ok := claims[auth.UserId].(string)
		if !ok {
			http.Error(w, "invalid claims", http.StatusBadRequest)
			return
		}
		if err := userId.From(claimsUserId); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rank, delta, updated, err := badges.GetMyRank(container, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rank == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		resp := GetBadgeRankResponse{
			Rank:    rank,
			Delta:   delta,
			Updated: updated,
		}

		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
