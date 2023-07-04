package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/snowmerak/ggeco/function/image"
	"github.com/snowmerak/ggeco/function/place"
	"github.com/snowmerak/ggeco/function/places"
	"github.com/snowmerak/ggeco/gen/bean"
	"github.com/snowmerak/ggeco/lib/maps"
	"github.com/snowmerak/ggeco/lib/sqlserver"
	"github.com/snowmerak/ggeco/lib/storage"
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

	http.HandleFunc("/api/place", place.Handler(container))
	http.HandleFunc("/api/search", places.Handler(container))
	http.HandleFunc("/api/image", image.Handler(container))
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
