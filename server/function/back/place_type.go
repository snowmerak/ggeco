package back

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"net/http"
)

type AddPlaceTypeToBadgeRequest struct {
	PlaceType string `json:"place_type"`
	BadgeId   string `json:"badge_id"`
}

func AddPlaceTypeToBadge(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		var r AddPlaceTypeToBadgeRequest
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&r); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		badgeId := sqlserver.UUID{}
		if err := badgeId.From(r.BadgeId); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := badges.AddPlaceTypeToBadgeId(container, r.PlaceType, badgeId); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		wr.WriteHeader(http.StatusCreated)
	}
}
