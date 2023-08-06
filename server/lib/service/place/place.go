package place

import (
	"context"
	"encoding/json"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"log"
	"time"
)

type GetPlaceRequest struct {
	PlaceId string `query:"place_id" required:"true"`
}

type GetPlaceResponse struct {
	ID         string                     `json:"id,omitempty"`
	Data       maps.SearchPlaceIdResponse `json:"Data,omitempty"`
	LastUpdate time.Time                  `json:"last_update,omitempty"`
}

func GetPlace(container sqlserver.Container, placeId string) (place GetPlaceResponse, err error) {
	sqlClient, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := sqlClient.Prepare("SELECT [data], [last_update] from [dbo].[Places] WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	cached := true
	row := stmt.QueryRow(placeId)
	if err := row.Err(); err != nil {
		cached = false
	}
	place.ID = placeId

	data := []byte(nil)
	if cached {
		func() {
			if err = row.Scan(&data, &place.LastUpdate); err != nil {
				cached = false
				return
			}

			if place.LastUpdate.AddDate(0, 6, 0).Before(time.Now()) {
				log.Println("cache expired")
				cached = false
				return
			}

			if err = json.Unmarshal(data, &place.Data); err != nil {
				log.Println(err)
				cached = false
				return
			}
		}()
	}

	if !cached {
		getOptions := []maps.SearchPlaceIdRequestOptional{func(request *maps.SearchPlaceIdRequest) *maps.SearchPlaceIdRequest {
			request.PlaceID = placeId
			request.Language = "ko"
			request.GetReviews = true
			request.GetOpeningHours = true
			request.GetPermanentlyClosed = true
			request.GetGeometryLocation = true
			request.GetPhotos = true
			request.GetPhone = true
			return request
		}}

		place.Data, err = maps.SearchPlaceId(context.TODO(), container, getOptions...)
		if err != nil {
			return
		}

		data, err = json.Marshal(place.Data)
		if err != nil {
			return
		}

		place.LastUpdate = time.Now()
		_, err = sqlClient.Exec("INSERT INTO [dbo].[Places] ([id], [data], [last_update]) VALUES (@P1, @P2, @P3)", place.ID, data, place.LastUpdate)
		if err != nil {
			return
		}
	}

	return
}
