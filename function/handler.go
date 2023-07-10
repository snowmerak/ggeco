package main

import (
	"context"
	maps2 "github.com/snowmerak/ggeco/lib/client/maps"
	sqlserver2 "github.com/snowmerak/ggeco/lib/client/sqlserver"
	storage2 "github.com/snowmerak/ggeco/lib/client/storage"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/snowmerak/ggeco/function/image"
	"github.com/snowmerak/ggeco/function/place"
	"github.com/snowmerak/ggeco/function/places"
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
	mapsClient, err := maps2.New(os.Getenv("GOOGLE_MAPS_API_KEY"), os.Getenv("GOOGLE_API_SIGNATURE"))
	if err != nil {
		panic(err)
	}
	maps2.PushClient(container, mapsClient)

	sqlServerHost := os.Getenv("SQL_SERVER_HOST")
	sqlServerPort, _ := strconv.ParseInt(os.Getenv("SQL_SERVER_PORT"), 10, 64)
	sqlServerUser := os.Getenv("SQL_SERVER_USER")
	sqlServerPassword := os.Getenv("SQL_SERVER_PASSWORD")
	sqlServerDatabase := os.Getenv("SQL_SERVER_DATABASE")
	sqlClient, err := sqlserver2.New(ctx, sqlServerHost, int(sqlServerPort), sqlServerUser, sqlServerPassword, sqlServerDatabase)
	if err != nil {
		panic(err)
	}
	sqlserver2.PushClient(container, sqlClient)

	imageClient, err := storage2.New(os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY"))
	if err != nil {
		panic(err)
	}
	storage2.PushClient(container, imageClient)

	http.HandleFunc("/api/place", place.Place(container))
	http.HandleFunc("/api/place/favorite/count", place.FavoriteCount(container))
	http.HandleFunc("/api/places", places.Handler(container))
	http.HandleFunc("/api/image", image.Handler(container))
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
