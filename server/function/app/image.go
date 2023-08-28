package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/imgmnger"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/auth"
	"log"
	"net/http"
)

const storageName = "image-silo"
const thumbnailSize = 256

type UploadImageRequest struct {
	Image string `file:"image" required:"true"`
	Size  int    `query:"size"`
}

type UploadImageResponse struct {
	OriginImageUrl    string `json:"origin_image_url"`
	ThumbnailImageUrl string `json:"thumbnail_image_url"`
}

func UploadImage(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println("UploadImage")
		imageManager, err := imgmnger.GetClient(container)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		claims, err := auth.GetJwtClaims(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := sqlserver.UUID{}
		claimsUserId, ok := claims[auth.UserId].(string)
		if !ok {
			log.Println(err)
			http.Error(w, "invalid claims", http.StatusInternalServerError)
			return
		}
		if err := userId.From(claimsUserId); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f, header, err := r.FormFile("image")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close()

		contentType := header.Header.Get("Content-Type")
		resp, err := imageManager.Upload(storageName, userId.String(), f, header.Filename, contentType, thumbnailSize)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(resp); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}
