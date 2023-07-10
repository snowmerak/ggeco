package place

import (
	"context"
	"github.com/snowmerak/ggeco/lib/client/maps"
)

func Search(container maps.Container, opts ...maps.SearchPlaceIdRequestOptional) (response maps.SearchPlaceIdResponse, err error) {
	return maps.SearchPlaceId(context.TODO(), container, opts...)
}
