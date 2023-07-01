package main

import (
	"github.com/snowmerak/ggeco/function/place"
	"github.com/snowmerak/ggeco/function/search"
	"github.com/snowmerak/ggeco/gen/bean"
	"github.com/snowmerak/ggeco/lib/maps"
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	container := bean.NewContainer()
	mapsClient, err := maps.New(os.Getenv("GOOGLE_MAPS_API_KEY"))
	if err != nil {
		panic(err)
	}
	maps.PushClient(container, mapsClient)

	http.HandleFunc("/api/place", place.Handler(container))
	http.HandleFunc("/api/search", search.Handler(container))
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
