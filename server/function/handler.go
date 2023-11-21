package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/snowmerak/ggeco/server/function/back"
	"github.com/snowmerak/ggeco/server/lib/client/imgmnger"

	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/ggeco/server/function/app"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/client/storage"
	"github.com/snowmerak/ggeco/server/lib/service/auth"

	"github.com/snowmerak/ggeco/server/gen/bean"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("panic: %v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	jwtSecretKey, err := base64.URLEncoding.DecodeString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		panic(err)
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

	blobStorageUrl := os.Getenv("IMAGE_MANAGER_APP_URL")
	if blobStorageUrl == "" {
		panic("IMAGE_MANAGER_APP_URL is empty")
	}
	imageManager := imgmnger.New(blobStorageUrl)
	imgmnger.PushClient(container, imageManager)

	router := httprouter.New()

	//router.GET("/api/place", api.Place(container))
	//router.GET("/api/places", api.Search(container))
	//
	//router.GET("/api/image", api.Image(container))
	//
	//router.GET("/api/course", api.GetCourse(container))
	//router.GET("/api/course/list", api.GetCourse(container))
	//router.POST("/api/course/name", api.UpdateCourseName(container))
	//router.POST("/api/course/review", api.UpdateCourseReview(container))
	//
	//router.GET("/api/badge", api.GetBadge(container))
	//router.GET("/api/badges", api.GetBadges(container))
	//router.GET("/api/badge/name", api.GetBadgeByName(container))
	//router.POST("/api/badge", api.AddBadge(container))
	//router.POST("/api/badge/name", api.UpdateBadgeName(container))
	//router.POST("/api/badge/summary", api.UpdateBadgeSummary(container))
	//router.POST("/api/badge/image", api.UpdateBadgeImage(container))
	//router.DELETE("/api/badge", api.DeleteBadge(container))
	//
	//router.POST("/api/badge/earned", api.AddEarnedBadge(container))
	//router.GET("/api/badge/earned", api.GetEarnedBadgesByUserId(container))
	//router.GET("/api/badge/earned/count", api.CountUsersEarnedBadge(container))
	//
	//router.GET("/api/course/place", api.GetCoursePlaces(container))
	//router.POST("/api/course/place", api.SetCoursePlaces(container))
	//
	//router.GET("/api/course/place/reviews", api.GetPlaceReviewsOfCourse(container))
	//router.POST("/api/course/place/review", api.CreatePlaceReview(container))
	//router.PATCH("/api/course/place/review", api.UpdatePlaceReview(container))
	//router.DELETE("/api/course/place/review", api.DeletePlaceReview(container))
	//
	//router.GET("/api/course/place/review/pic", api.GetPlaceReviewPictures(container))
	//router.POST("/api/course/place/review/pic", api.SetPlaceReviewPictures(container))
	//
	//router.GET("/api/course/badges", api.GetCourseBadges(container))
	//router.POST("/api/course/badges", api.SetCourseBadges(container))
	//
	//router.GET("/api/user", api.GetUser(container))
	//router.POST("/api/user", api.AddUser(container))
	//router.PATCH("/api/user/nickname", api.UpdateUserNickname(container))
	//router.PATCH("/api/user/last_signin", api.UpdateUserLastSigninDate(container))
	//router.DELETE("/api/user", api.DeleteUser(container))
	//
	//router.GET("/api/user/naver", api.GetNaverUser(container))
	//router.GET("/api/user/naver/id", api.GetNaverUserByNaverId(container))
	//router.POST("/api/user/naver", api.AddNaverUser(container))
	//router.PATCH("/api/user/naver", api.UpdateNaverUser(container))
	//router.DELETE("/api/user/naver", api.DeleteNaverUser(container))
	//
	//router.GET("/api/user/kakao", api.GetKakaoUser(container))
	//router.GET("/api/user/kakao/id", api.GetKakaoUserByKakaoId(container))
	//router.POST("/api/user/kakao", api.AddKakaoUser(container))
	//router.PATCH("/api/user/kakao", api.UpdateKakaoUser(container))
	//router.DELETE("/api/user/kakao", api.DeleteKakaoUser(container))
	//
	//router.GET("/api/place/favorite/user", api.GetFavoritePlacesByUserId(container))
	//router.GET("/api/place/favorite/count", api.CountFavoritePlace(container))
	//router.POST("/api/place/favorite", api.AddFavoritePlace(container))
	//router.DELETE("/api/place/favorite", api.DeleteFavoritePlace(container))
	//
	//router.GET("/api/course/favorite/user", api.GetFavoriteCoursesByUserId(container))
	//router.GET("/api/course/favorite/count", api.CountFavoriteCourse(container))
	//router.POST("/api/course/favorite", api.AddFavoriteCourse(container))
	//router.DELETE("/api/course/favorite", api.DeleteFavoriteCourse(container))

	//////////////////////////////////////////////////////////////////////////////////////////////////////

	router.POST("/app/auth/signin", app.SignIn(container))
	router.POST("/app/auth/refresh", app.Refresh(container))

	router.GET("/app/badge/list", app.GetBadges(container))
	router.GET("/app/badge/searchables", app.GetSearchableBadges(container))
	router.GET("/app/badge/earned", app.GetEarnedBadges(container))
	router.GET("/app/badge", app.GetBadgeInfo(container))
	router.GET("/app/badge/rank", app.GetBadgeRank(container))

	router.GET("/app/course/popular", app.GetPopularCourseOfBadge(container))
	router.GET("/app/course/recent", app.GetRecentCourses(container))
	router.GET("/app/course/recommend", app.GetRecommendedCourses(container))
	router.GET("/app/course/search", app.FindCoursesBySearchPlace(container))
	router.GET("/app/course/my", app.GetMyCourses(container))
	router.GET("/app/course", app.GetCourseInfo(container))
	router.POST("/app/course", app.AddCourse(container))
	router.POST("/app/course/edit", app.UpdateCourse(container))
	router.GET("/app/course/favorite", app.GetFavoriteCourses(container))
	router.POST("/app/course/favorite", app.AddFavoriteCourse(container))
	router.DELETE("/app/course/favorite", app.RemoveFavoriteCourse(container))
	router.GET("/app/course/favorite/check", app.IsFavoriteCourse(container))

	router.GET("/app/place/search", app.SearchPlaces(container))
	router.GET("/app/place", app.GetPlaceInfo(container))
	router.GET("/app/place/favorite", app.GetFavoritePlaces(container))
	router.POST("/app/place/favorite", app.AddFavoritePlace(container))
	router.DELETE("/app/place/favorite", app.RemoveFavoritePlace(container))
	router.GET("/app/place/favorite/check", app.IsFavoritePlace(container))

	router.GET("/app/user", app.GetProfile(container))
	router.POST("/app/user/nickname", app.UpdateNickname(container))
	router.POST("/app/user/badge", app.UpdateBadge(container))

	router.POST("/back/badge", back.AddBadge(container))
	router.POST("/back/placetype", back.AddPlaceTypeToBadge(container))

	// router.ServeFiles("/app/swagger/*filepath", http.Dir("./swagger"))

	listenFullAddr := fmt.Sprintf("https://127.0.0.1%s/", listenAddr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
		reqCtx := auth.WithJwtSecretKey(r.Context(), jwtSecretKey)
		r = r.WithContext(reqCtx)
		authorization := r.Header.Get("Authorization")
		if strings.HasPrefix(authorization, "Bearer ") {
			authorization = strings.TrimPrefix(authorization, "Bearer ")
			claims, err := auth.ValidateUserToken(jwtSecretKey, authorization)
			if err == nil {
				reqCtx = auth.WithJwtClaims(reqCtx, claims)
				r = r.WithContext(reqCtx)
			}
		}
		router.ServeHTTP(wr, r)
	})

	log.Printf("About to listen on %s. Go to %s", listenAddr, listenFullAddr)
	log.Fatal().Msg(http.ListenAndServe(listenAddr, mux).Error())
}
