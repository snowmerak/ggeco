package back

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/storage"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"io"
	"net/http"
	"strconv"
)

type Badge struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Summary       string `json:"summary,omitempty"`
	ActiveImage   string `json:"active_image,omitempty"`
	InactiveImage string `json:"inactive_image,omitempty"`
	SelectedImage string `json:"selected_image,omitempty"`
}

type AddBadgeRequest struct {
	Name        string `form:"name" required:"true"`
	Description string `form:"description" required:"true"`
	Image       string `file:"image" required:"true"`
}

type AddBadgeResponse struct {
	Badge Badge `json:"badge"`
}

func AddBadge(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}

		description := r.FormValue("description")
		if description == "" {
			http.Error(w, "description is required", http.StatusBadRequest)
			return
		}

		searchableValue := r.FormValue("searchable")
		if searchableValue == "" {
			searchableValue = "false"
		}
		searchable, err := strconv.ParseBool(searchableValue)
		if err != nil {
			searchable = false
		}

		urls := []string{"", "", ""}
		for i, n := range []string{"active_image", "inactive_image", "selected_image"} {
			f, header, err := r.FormFile(n)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			data, err := io.ReadAll(f)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			url, err := storage.UploadImage(container, "image-silo", header.Filename, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			urls[i] = url
		}

		badgeId, err := badges.Add(container, name, description, urls[0], urls[1], urls[2], searchable)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		badge := Badge{
			Id:            badgeId.String(),
			Name:          name,
			Summary:       description,
			ActiveImage:   urls[0],
			InactiveImage: urls[1],
			SelectedImage: urls[2],
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(badge); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
