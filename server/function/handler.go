package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/function/api"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/client/storage"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/snowmerak/ggeco/server/gen/bean"
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

	router.GET("/api/badge", api.GetBadge(container))
	router.GET("/api/badges", api.GetBadges(container))
	router.GET("/api/badge/name", api.GetBadgeByName(container))
	router.POST("/api/badge", api.AddBadge(container))
	router.POST("/api/badge/name", api.UpdateBadgeName(container))
	router.POST("/api/badge/summary", api.UpdateBadgeSummary(container))
	router.POST("/api/badge/image", api.UpdateBadgeImage(container))
	router.DELETE("/api/badge", api.DeleteBadge(container))

	router.POST("/api/badge/earned", api.AddEarnedBadge(container))
	router.GET("/api/badge/earned", api.GetEarnedBadgesByUserId(container))
	router.GET("/api/badge/earned/count", api.CountUsersEarnedBadge(container))

	router.GET("/api/course/place", api.GetCoursePlaces(container))
	router.POST("/api/course/place", api.SetCoursePlaces(container))

	router.GET("/api/course/place/reviews", api.GetPlaceReviewsOfCourse(container))
	router.POST("/api/course/place/review", api.CreatePlaceReview(container))
	router.PATCH("/api/course/place/review", api.UpdatePlaceReview(container))
	router.DELETE("/api/course/place/review", api.DeletePlaceReview(container))

	router.GET("/api/course/place/review/pic", api.GetPlaceReviewPictures(container))
	router.POST("/api/course/place/review/pic", api.SetPlaceReviewPictures(container))

	router.GET("/api/course/badges", api.GetCourseBadges(container))
	router.POST("/api/course/badges", api.SetCourseBadges(container))

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
