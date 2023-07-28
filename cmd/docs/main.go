package main

import (
	"fmt"
	"github.com/snowmerak/ggeco/lib/client/maps"
	"github.com/snowmerak/ggeco/lib/service/place"
	"github.com/swaggest/openapi-go/openapi3"
	"net/http"
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

	placesGetOp := openapi3.Operation{}
	placesGetOp.WithDescription("Search Places from google maps api").
		WithSummary("Search Places").
		WithParameters(openapi3.ParameterOrRef{
			Parameter: &openapi3.Parameter{
				Name:     "query",
				In:       "query",
				Required: wrap(true),
				Schema: &openapi3.SchemaOrRef{
					Schema: &openapi3.Schema{
						Type: wrap(openapi3.SchemaTypeString),
					},
				},
				Description: wrap("Search query"),
			},
		}, openapi3.ParameterOrRef{
			Parameter: &openapi3.Parameter{
				Name: "lang",
				In:   "query",
				Schema: &openapi3.SchemaOrRef{
					Schema: &openapi3.Schema{
						Type: wrap(openapi3.SchemaTypeString),
					},
				},
			},
		}, openapi3.ParameterOrRef{
			Parameter: &openapi3.Parameter{
				Name: "latitude",
				In:   "query",
				Schema: &openapi3.SchemaOrRef{
					Schema: &openapi3.Schema{
						Type: wrap(openapi3.SchemaTypeNumber),
					},
				},
			},
		}, openapi3.ParameterOrRef{
			Parameter: &openapi3.Parameter{
				Name: "longitude",
				In:   "query",
				Schema: &openapi3.SchemaOrRef{
					Schema: &openapi3.Schema{
						Type: wrap(openapi3.SchemaTypeNumber),
					},
				},
			},
		}, openapi3.ParameterOrRef{
			Parameter: &openapi3.Parameter{
				Name: "radius",
				In:   "query",
				Schema: &openapi3.SchemaOrRef{
					Schema: &openapi3.Schema{
						Type: wrap(openapi3.SchemaTypeNumber),
					},
				},
			},
		})
	reflector.SetJSONResponse(&placesGetOp, maps.SearchTextResponse{}, http.StatusOK)
	reflector.Spec.AddOperation(http.MethodGet, "/places", placesGetOp)

	placeGetOp := openapi3.Operation{}
	placeGetOp.WithDescription("Get Place details from Google Maps API or cached").
		WithSummary("Get place details").
		WithParameters(openapi3.ParameterOrRef{
			Parameter: &openapi3.Parameter{
				Name:     "place_id",
				In:       "query",
				Required: wrap(true),
				Schema: &openapi3.SchemaOrRef{
					Schema: &openapi3.Schema{
						Type: wrap(openapi3.SchemaTypeString),
					},
				},
			},
		}).
		WithResponses(openapi3.Responses{
			MapOfResponseOrRefValues: map[string]openapi3.ResponseOrRef{
				"500": {
					Response: &openapi3.Response{
						Description: "Internal Server Error or Google Maps API Error.",
					},
				},
			},
		})
	reflector.SetJSONResponse(&placeGetOp, place.Information{}, http.StatusOK)
	reflector.Spec.AddOperation(http.MethodGet, "/place", placeGetOp)

	imageGetOp := openapi3.Operation{}
	imageGetOp.WithDescription("Get Path of the Image").
		WithSummary("Get Image Path").
		WithParameters(openapi3.ParameterOrRef{
			Parameter: &openapi3.Parameter{
				Name:     "name",
				In:       "query",
				Required: wrap(true),
				Schema: &openapi3.SchemaOrRef{
					Schema: &openapi3.Schema{
						Type: wrap(openapi3.SchemaTypeString),
					},
				},
			},
		}).
		WithResponses(openapi3.Responses{
			MapOfResponseOrRefValues: map[string]openapi3.ResponseOrRef{
				"500": {
					Response: &openapi3.Response{
						Description: "Internal Server Error. Azure Blob Storage error.",
					},
				},
				"200": {
					Response: &openapi3.Response{
						Description: "A shared access signature (SAS) URI providing write access to the blob.",
						Content: map[string]openapi3.MediaType{
							"text/plain": {
								Schema: &openapi3.SchemaOrRef{
									Schema: &openapi3.Schema{
										Type: wrap(openapi3.SchemaTypeString),
									},
								},
							},
						},
					},
				},
			},
		})
	reflector.Spec.AddOperation(http.MethodGet, "/image", imageGetOp)

	value, err := reflector.Spec.MarshalYAML()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(value))
}
