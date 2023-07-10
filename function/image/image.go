package image

import (
	"github.com/snowmerak/ggeco/gen/bean"
	"github.com/snowmerak/ggeco/lib/client/storage"
	"net/http"
)

func Handler(container *bean.Container) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("name")
		url, err := storage.GetSASURL(container, "image-silo", filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(url))
	}
}
