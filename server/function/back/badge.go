package back

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/storage"
	"github.com/snowmerak/ggeco/server/lib/service/badges"
	"io"
	"net/http"
)

type Badge struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
	Image   string `json:"image,omitempty"`
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

		f, header, err := r.FormFile("image")
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

		badgeId, err := badges.Add(container, name, description, url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		badge := Badge{
			Id:      badgeId.String(),
			Name:    name,
			Summary: description,
			Image:   url,
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(badge); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
