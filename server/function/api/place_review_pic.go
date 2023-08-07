package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/service/place"
	"net/http"
)

func GetPlaceReviewPictures(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		id := r.URL.Query().Get("review_id")
		uuid, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		pics, err := place.GetReviewPictures(container, uuid)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := place.GetReviewPicturesResponse{}
		for _, pic := range pics {
			resp.Pictures = append(resp.Pictures, pic.PictureUrl)
			resp.Thumbnails = append(resp.Thumbnails, pic.ThumbnailUrl)
		}

		encoder := json.NewEncoder(wr)
		if err := encoder.Encode(resp); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func SetPlaceReviewPictures(container bean.Container) httprouter.Handle {
	return func(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		req := place.SetReviewPicturesRequest{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		reviewUUID, err := base64.URLEncoding.DecodeString(req.ReviewId)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := place.SetReviewPictures(container, reviewUUID, req.PictureUrls, req.ThumbnailUrls); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}

		wr.WriteHeader(http.StatusOK)
	}
}
