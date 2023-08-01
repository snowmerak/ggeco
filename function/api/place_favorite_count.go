package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/gen/bean"
	"net/http"
)

func FavoriteCount(container bean.Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("place favorite count"))
	}
}
