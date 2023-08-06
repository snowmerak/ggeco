package image

import (
	"context"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/storage"
)

type GetImageURLRequest struct {
	Filename string `query:"name,required"`
}

func GetImageURL(ctx context.Context, container bean.Container, filename string) (url string, err error) {
	url, err = storage.GetSASURL(container, "image-silo", filename)
	if err != nil {
		return
	}

	return
}
