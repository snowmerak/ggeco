package place

import (
	"github.com/snowmerak/ggeco/gen/bean"
	"net/http"
)

func FavoriteCount(container *bean.Container) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("place favorite count"))
	}
}
