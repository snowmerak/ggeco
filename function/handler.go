package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/function/api"
	"github.com/snowmerak/ggeco/lib/client/maps"
	"github.com/snowmerak/ggeco/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/lib/client/storage"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/snowmerak/ggeco/gen/bean"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	container := bean.NewContainer()
	mapsClient, err := maps.New(os.Getenv("GOOGLE_MAPS_API_KEY"), os.Getenv("GOOGLE_API_SIGNATURE"))
	if err != nil {
		panic(err)
	}
	maps.PushClient(container, mapsClient)

	sqlServerHost := os.Getenv("SQL_SERVER_HOST")
	sqlServerPort, _ := strconv.ParseInt(os.Getenv("SQL_SERVER_PORT"), 10, 64)
	sqlServerUser := os.Getenv("SQL_SERVER_USER")
	sqlServerPassword := os.Getenv("SQL_SERVER_PASSWORD")
	sqlServerDatabase := os.Getenv("SQL_SERVER_DATABASE")
	sqlClient, err := sqlserver.New(ctx, sqlServerHost, int(sqlServerPort), sqlServerUser, sqlServerPassword, sqlServerDatabase)
	if err != nil {
		panic(err)
	}
	sqlserver.PushClient(container, sqlClient)

	imageClient, err := storage.New(os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY"))
	if err != nil {
		panic(err)
	}
	storage.PushClient(container, imageClient)

	router := httprouter.New()

	router.GET("/api/place", api.Place(container))
	router.GET("/api/place/favorite/count", api.FavoriteCount(container))
	router.GET("/api/places", api.Search(container))

	router.GET("/api/image", api.Image(container))

	router.GET("/api/course", api.GetCourse(container))
	router.GET("/api/course/list", api.GetCourse(container))
	router.POST("/api/course/name", api.UpdateCourseName(container))
	router.POST("/api/course/review", api.UpdateCourseReview(container))

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
