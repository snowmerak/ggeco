package main

import (
	"github.com/snowmerak/ggeco/lib/client/maps"
	"github.com/snowmerak/ggeco/lib/service/image"
	"github.com/snowmerak/ggeco/lib/service/place"
	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
	"net/http"
	"os"
)

func wrap[T any](t T) *T {
	return &t
}

func main() {
	reflector := openapi3.Reflector{
		Spec: &openapi3.Spec{
			Openapi: "3.0.0",
		},
	}

	author := "snowmerak"
	email := "snowmerak@outlook.com"

	reflector.Spec.Info.
		WithTitle("ggeco").
		WithDescription("The ggeco API Documentation").
		WithVersion("1.0.0").
		WithContact(openapi3.Contact{
			Name:  &author,
			Email: &email,
		})

	placesGetOp, err := reflector.NewOperationContext(http.MethodGet, "/places")
	if err != nil {
		panic(err)
	}
	placesGetOp.SetDescription("Search Places from google maps api")
	placesGetOp.SetSummary("Search Places")
	placesGetOp.AddReqStructure(place.SearchTextRequest{})
	placesGetOp.AddRespStructure([]maps.SearchTextResponse{})
	placesGetOp.AddRespStructure(nil, openapi.ContentOption(func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error or Google Maps API Error."
	}))
	reflector.AddOperation(placesGetOp)

	placeGetOp, err := reflector.NewOperationContext(http.MethodGet, "/place")
	if err != nil {
		panic(err)
	}
	placeGetOp.SetDescription("Get Place details from Google Maps API or cached")
	placeGetOp.SetSummary("Get place details")
	placeGetOp.AddReqStructure(place.GetPlaceRequest{})
	placeGetOp.AddRespStructure(place.GetPlaceResponse{})
	placeGetOp.AddRespStructure(nil, openapi.ContentOption(func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error or Google Maps API Error."
	}))
	reflector.AddOperation(placeGetOp)

	imageGetOp, err := reflector.NewOperationContext(http.MethodGet, "/image")
	if err != nil {
		panic(err)
	}
	imageGetOp.SetDescription("Get Path of the Image")
	imageGetOp.SetSummary("Get Image Path")
	imageGetOp.AddReqStructure(image.GetImageURLRequest{})
	imageGetOp.AddRespStructure("", func(cu *openapi.ContentUnit) {
		cu.ContentType = "text/plain"
	})
	imageGetOp.AddRespStructure(nil, openapi.ContentOption(func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error. Azure Blob Storage error."
	}))
	reflector.AddOperation(imageGetOp)

	value, err := reflector.Spec.MarshalYAML()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./doc/swagger.yaml")
	if err != nil {
		panic(err)
	}

	_, err = f.Write(value)
	if err != nil {
		panic(err)
	}
}
