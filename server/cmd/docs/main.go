package main

import (
	"github.com/snowmerak/ggeco/server/function/app"
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
		WithVersion("1.1.0").
		WithContact(openapi3.Contact{
			Name:  &author,
			Email: &email,
		})

	//placesGetOp, err := reflector.NewOperationContext(http.MethodGet, "/api/places")
	//if err != nil {
	//	panic(err)
	//}
	//placesGetOp.SetDescription("Search Places from google maps api")
	//placesGetOp.SetSummary("Search Places")
	//placesGetOp.AddReqStructure(place.SearchTextRequest{})
	//placesGetOp.AddRespStructure([]maps.SearchTextResponse{})
	//placesGetOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error or Google Maps API Error."
	//})
	//reflector.AddOperation(placesGetOp)
	//
	//placeGetOp, err := reflector.NewOperationContext(http.MethodGet, "/api/place")
	//if err != nil {
	//	panic(err)
	//}
	//placeGetOp.SetDescription("Get Place details from Google Maps API or cached")
	//placeGetOp.SetSummary("Get place details")
	//placeGetOp.AddReqStructure(place.GetPlaceRequest{})
	//placeGetOp.AddRespStructure(place.GetPlaceResponse{})
	//placeGetOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error or Google Maps API Error."
	//})
	//reflector.AddOperation(placeGetOp)
	//
	//imageGetOp, err := reflector.NewOperationContext(http.MethodGet, "/api/image")
	//if err != nil {
	//	panic(err)
	//}
	//imageGetOp.SetDescription("Get Path of the Image")
	//imageGetOp.SetSummary("Get Image Path")
	//imageGetOp.AddReqStructure(image.GetImageURLRequest{})
	//imageGetOp.AddRespStructure("", func(cu *openapi.ContentUnit) {
	//	cu.ContentType = "text/plain"
	//})
	//imageGetOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error. Azure Blob Storage error."
	//})
	//reflector.AddOperation(imageGetOp)
	//
	//courseGetOp, err := reflector.NewOperationContext(http.MethodGet, "/api/course")
	//if err != nil {
	//	panic(err)
	//}
	//courseGetOp.SetDescription("Get Course data by course id")
	//courseGetOp.SetSummary("Get Course data")
	//courseGetOp.AddReqStructure(courses.GetCourseRequest{})
	//courseGetOp.AddRespStructure(courses.Course{})
	//reflector.AddOperation(courseGetOp)
	//
	//courseListGetOp, err := reflector.NewOperationContext(http.MethodGet, "/api/course/list")
	//if err != nil {
	//	panic(err)
	//}
	//courseListGetOp.SetDescription("Get Course data by author id or course name")
	//courseListGetOp.SetSummary("Get Course List")
	//courseListGetOp.AddReqStructure(courses.GetCourseListRequest{})
	//courseListGetOp.AddRespStructure([]courses.Course{})
	//reflector.AddOperation(courseListGetOp)
	//
	//updateCourseNameOp, err := reflector.NewOperationContext(http.MethodPost, "/api/course/name")
	//if err != nil {
	//	panic(err)
	//}
	//updateCourseNameOp.SetDescription("Update Course Name")
	//updateCourseNameOp.SetSummary("Update Course Name")
	//updateCourseNameOp.AddReqStructure(courses.UpdateCourseNameRequest{})
	//updateCourseNameOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusOK
	//})
	//updateCourseNameOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error. Azure SQL Database error."
	//	cu.Format = "text/plain"
	//})
	//reflector.AddOperation(updateCourseNameOp)
	//
	//updateCourseReviewOp, err := reflector.NewOperationContext(http.MethodPost, "/api/course/review")
	//if err != nil {
	//	panic(err)
	//}
	//updateCourseReviewOp.SetDescription("Update Course Review")
	//updateCourseReviewOp.SetSummary("Update Course Review")
	//updateCourseReviewOp.AddReqStructure(courses.UpdateCourseReviewRequest{})
	//updateCourseReviewOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusOK
	//})
	//updateCourseReviewOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error. Azure SQL Database error."
	//	cu.Format = "text/plain"
	//})
	//reflector.AddOperation(updateCourseReviewOp)
	//
	//getBadgeOp, err := reflector.NewOperationContext(http.MethodGet, "/api/badge")
	//if err != nil {
	//	panic(err)
	//}
	//getBadgeOp.SetDescription("Get Badge data by badge id")
	//getBadgeOp.SetSummary("Get Badge data")
	//getBadgeOp.AddReqStructure(badges.GetBadgeRequest{})
	//getBadgeOp.AddRespStructure(badges.GetBadgeResponse{})
	//reflector.AddOperation(getBadgeOp)
	//
	//getBadgeByNameOp, err := reflector.NewOperationContext(http.MethodGet, "/api/badge/name")
	//if err != nil {
	//	panic(err)
	//}
	//getBadgeByNameOp.SetDescription("Get Badge data by badge name")
	//getBadgeByNameOp.SetSummary("Get Badge data")
	//getBadgeByNameOp.AddReqStructure(badges.GetBadgeByNameRequest{})
	//getBadgeByNameOp.AddRespStructure([]badges.GetBadgeResponse{})
	//reflector.AddOperation(getBadgeByNameOp)
	//
	//getBadgeListOp, err := reflector.NewOperationContext(http.MethodGet, "/api/badges")
	//if err != nil {
	//	panic(err)
	//}
	//getBadgeListOp.SetDescription("Get All Badge Data")
	//getBadgeListOp.SetSummary("Get Badge List")
	//getBadgeListOp.AddRespStructure([]badges.GetBadgeResponse{})
	//reflector.AddOperation(getBadgeListOp)
	//
	//addBadgeOp, err := reflector.NewOperationContext(http.MethodPost, "/api/badge")
	//if err != nil {
	//	panic(err)
	//}
	//addBadgeOp.SetDescription("Add Badge")
	//addBadgeOp.SetSummary("Add Badge")
	//addBadgeOp.AddReqStructure(badges.AddBadgeRequest{})
	//addBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusOK
	//})
	//addBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error. Azure SQL Database error."
	//	cu.Format = "text/plain"
	//})
	//reflector.AddOperation(addBadgeOp)
	//
	//updateBadgeNameOp, err := reflector.NewOperationContext(http.MethodPost, "/api/badge/name")
	//if err != nil {
	//	panic(err)
	//}
	//updateBadgeNameOp.SetDescription("Update Badge Name")
	//updateBadgeNameOp.SetSummary("Update Badge Name")
	//updateBadgeNameOp.AddReqStructure(badges.UpdateBadgeNameRequest{})
	//updateBadgeNameOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusOK
	//})
	//updateBadgeNameOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error. Azure SQL Database error."
	//	cu.Format = "text/plain"
	//})
	//reflector.AddOperation(updateBadgeNameOp)
	//
	//updateBadgeSummaryOp, err := reflector.NewOperationContext(http.MethodPost, "/api/badge/summary")
	//if err != nil {
	//	panic(err)
	//}
	//updateBadgeSummaryOp.SetDescription("Update Badge Summary")
	//updateBadgeSummaryOp.SetSummary("Update Badge Summary")
	//updateBadgeSummaryOp.AddReqStructure(badges.UpdateBadgeSummaryRequest{})
	//updateBadgeSummaryOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusOK
	//})
	//updateBadgeSummaryOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error. Azure SQL Database error."
	//	cu.Format = "text/plain"
	//})
	//reflector.AddOperation(updateBadgeSummaryOp)
	//
	//updateBadgeImageOp, err := reflector.NewOperationContext(http.MethodPost, "/api/badge/image")
	//if err != nil {
	//	panic(err)
	//}
	//updateBadgeImageOp.SetDescription("Update Badge Image")
	//updateBadgeImageOp.SetSummary("Update Badge Image")
	//updateBadgeImageOp.AddReqStructure(badges.UpdateBadgeImageRequest{})
	//updateBadgeImageOp.AddRespStructure(nil, openapi.ContentOption(func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusOK
	//}))
	//updateBadgeImageOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error. Azure SQL Database error."
	//	cu.Format = "text/plain"
	//})
	//reflector.AddOperation(updateBadgeImageOp)
	//
	//deleteBadgeOp, err := reflector.NewOperationContext(http.MethodDelete, "/api/badge")
	//if err != nil {
	//	panic(err)
	//}
	//deleteBadgeOp.SetDescription("Delete Badge")
	//deleteBadgeOp.SetSummary("Delete Badge")
	//deleteBadgeOp.AddReqStructure(badges.DeleteBadgeRequest{})
	//deleteBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusOK
	//})
	//deleteBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusInternalServerError
	//	cu.Description = "Internal Server Error. Azure SQL Database error."
	//	cu.Format = "text/plain"
	//})
	//reflector.AddOperation(deleteBadgeOp)
	//
	//getEarnedBadgesByUserIdOp, err := reflector.NewOperationContext(http.MethodGet, "/api/badges/earned")
	//if err != nil {
	//	panic(err)
	//}
	//getEarnedBadgesByUserIdOp.SetDescription("Get Earned Badges By User Id")
	//getEarnedBadgesByUserIdOp.SetSummary("Get Earned Badges By User Id")
	//getEarnedBadgesByUserIdOp.AddReqStructure(badges.GetEarnedBadgesRequest{})
	//getEarnedBadgesByUserIdOp.AddRespStructure([]badges.GetEarnedBadgeResponse{})
	//reflector.AddOperation(getEarnedBadgesByUserIdOp)
	//
	//countUsersEarnedBadgeOp, err := reflector.NewOperationContext(http.MethodGet, "/api/badges/earned/count")
	//if err != nil {
	//	panic(err)
	//}
	//countUsersEarnedBadgeOp.SetDescription("Count Users Earned Badge")
	//countUsersEarnedBadgeOp.SetSummary("Count Users Earned Badge")
	//countUsersEarnedBadgeOp.AddReqStructure(badges.CountUsersEarnedBadgeRequest{})
	//countUsersEarnedBadgeOp.AddRespStructure(badges.CountUsersEarnedBadgeResponse{})
	//reflector.AddOperation(countUsersEarnedBadgeOp)
	//
	//addEarnedBadgeOp, err := reflector.NewOperationContext(http.MethodPost, "/api/badges/earned")
	//if err != nil {
	//	panic(err)
	//}
	//addEarnedBadgeOp.SetDescription("Add Earned Badge")
	//addEarnedBadgeOp.SetSummary("Add Earned Badge")
	//addEarnedBadgeOp.AddReqStructure(badges.AddEarnedBadgeRequest{})
	//addEarnedBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
	//	cu.HTTPStatus = http.StatusOK
	//})
	//reflector.AddOperation(addEarnedBadgeOp)

	appSigninOp, err := reflector.NewOperationContext(http.MethodPost, "https://ggeco-func.azurewebsites.net/app/auth/signin")
	if err != nil {
		panic(err)
	}
	appSigninOp.SetDescription("Sign in with OAuth2")
	appSigninOp.SetSummary("Sign in with OAuth2")
	appSigninOp.AddReqStructure(app.SignInRequest{})
	appSigninOp.AddRespStructure(app.SignInResponse{})
	appSigninOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appSigninOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appSigninOp); err != nil {
		panic(err)
	}

	appRefreshOp, err := reflector.NewOperationContext(http.MethodPost, "https://ggeco-func.azurewebsites.net/app/auth/refresh")
	if err != nil {
		panic(err)
	}
	appRefreshOp.SetDescription("Refresh Access Token")
	appRefreshOp.SetSummary("Refresh Access Token")
	appRefreshOp.AddReqStructure(app.RefreshRequest{})
	appRefreshOp.AddRespStructure(app.RefreshResponse{})
	appRefreshOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appRefreshOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appRefreshOp); err != nil {
		panic(err)
	}

	appGetBadgesOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/badge/list")
	if err != nil {
		panic(err)
	}
	appGetBadgesOp.SetDescription("Get Badge List")
	appGetBadgesOp.SetSummary("Get Badge List")
	appGetBadgesOp.AddReqStructure(app.GetBadgesRequest{})
	appGetBadgesOp.AddRespStructure(app.GetBadgesResponse{})
	appGetBadgesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetBadgesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetBadgesOp); err != nil {
		panic(err)
	}

	appGetEarnedBadgesOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/badge/earned")
	if err != nil {
		panic(err)
	}
	appGetEarnedBadgesOp.SetDescription("Get Earned Badges")
	appGetEarnedBadgesOp.SetSummary("Get Earned Badges")
	appGetEarnedBadgesOp.AddReqStructure(app.GetEarnedBadgesRequest{})
	appGetEarnedBadgesOp.AddRespStructure(app.GetEarnedBadgesResponse{})
	appGetEarnedBadgesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetEarnedBadgesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetEarnedBadgesOp); err != nil {
		panic(err)
	}

	appGetBadgeOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/badge")
	if err != nil {
		panic(err)
	}
	appGetBadgeOp.SetDescription("Get Badge Info")
	appGetBadgeOp.SetSummary("Get Badge Info")
	appGetBadgeOp.AddReqStructure(app.GetBadgeInfoRequest{})
	appGetBadgeOp.AddRespStructure(app.GetBadgeInfoResponse{})
	appGetBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetBadgeOp); err != nil {
		panic(err)
	}

	appGetPopularCourseOfBadgeOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/course/popular")
	if err != nil {
		panic(err)
	}
	appGetPopularCourseOfBadgeOp.SetDescription("Get Popular Course of Badge")
	appGetPopularCourseOfBadgeOp.SetSummary("Get Popular Course of Badge")
	appGetPopularCourseOfBadgeOp.AddReqStructure(app.GetPopularCourseOfBadgeRequest{})
	appGetPopularCourseOfBadgeOp.AddRespStructure(app.GetPopularCourseOfBadgeResponse{})
	appGetPopularCourseOfBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetPopularCourseOfBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetPopularCourseOfBadgeOp); err != nil {
		panic(err)
	}

	appGetRecentCoursesOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/course/recent")
	if err != nil {
		panic(err)
	}
	appGetRecentCoursesOp.SetDescription("Get Recent Courses")
	appGetRecentCoursesOp.SetSummary("Get Recent Courses")
	appGetRecentCoursesOp.AddReqStructure(app.GetRecentCoursesRequest{})
	appGetRecentCoursesOp.AddRespStructure(app.GetRecentCoursesResponse{})
	appGetRecentCoursesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetRecentCoursesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetRecentCoursesOp); err != nil {
		panic(err)
	}

	appFindCoursesBySearchPlaceOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/course/search")
	if err != nil {
		panic(err)
	}
	appFindCoursesBySearchPlaceOp.SetDescription("Find Courses by Search Place")
	appFindCoursesBySearchPlaceOp.SetSummary("Find Courses by Search Place")
	appFindCoursesBySearchPlaceOp.AddReqStructure(app.FindCoursesBySearchPlaceRequest{})
	appFindCoursesBySearchPlaceOp.AddRespStructure(app.FindCoursesBySearchPlaceResponse{})
	appFindCoursesBySearchPlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appFindCoursesBySearchPlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appFindCoursesBySearchPlaceOp); err != nil {
		panic(err)
	}

	getMyCoursesOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/course/my")
	if err != nil {
		panic(err)
	}
	getMyCoursesOp.SetDescription("Get My Courses")
	getMyCoursesOp.SetSummary("Get My Courses")
	getMyCoursesOp.AddRespStructure(app.GetMyCoursesResponse{})
	getMyCoursesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	getMyCoursesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})

	appGetCourseOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/course")
	if err != nil {
		panic(err)
	}
	appGetCourseOp.SetDescription("Get Course Info")
	appGetCourseOp.SetSummary("Get Course Info")
	appGetCourseOp.AddReqStructure(app.GetCourseInfoRequest{})
	appGetCourseOp.AddRespStructure(app.GetCourseInfoResponse{})
	appGetCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetCourseOp); err != nil {
		panic(err)
	}

	appAddCourseOp, err := reflector.NewOperationContext(http.MethodPost, "https://ggeco-func.azurewebsites.net/app/course")
	if err != nil {
		panic(err)
	}
	appAddCourseOp.SetDescription("Set Course Data")
	appAddCourseOp.SetSummary("Set Course Data")
	appAddCourseOp.AddReqStructure(app.SetCourseRequest{})
	appAddCourseOp.AddRespStructure(app.SetCourseResponse{})
	appAddCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appAddCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appAddCourseOp); err != nil {
		panic(err)
	}

	appUpdateCourseOp, err := reflector.NewOperationContext(http.MethodPost, "https://ggeco-func.azurewebsites.net/app/course/edit")
	if err != nil {
		panic(err)
	}
	appUpdateCourseOp.SetDescription("Update Course Data")
	appUpdateCourseOp.SetSummary("Update Course Data")
	appUpdateCourseOp.AddReqStructure(app.UpdateCourseRequest{})
	appUpdateCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusOK
		cu.Description = "Update Course Data Success."
	})
	appUpdateCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appUpdateCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appUpdateCourseOp); err != nil {
		panic(err)
	}

	appGetFavoriteCoursesOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/course/favorite")
	if err != nil {
		panic(err)
	}
	appGetFavoriteCoursesOp.SetDescription("Get Favorite Courses")
	appGetFavoriteCoursesOp.SetSummary("Get Favorite Courses")
	appGetFavoriteCoursesOp.AddRespStructure(app.GetFavoriteCoursesResponse{})
	appGetFavoriteCoursesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetFavoriteCoursesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetFavoriteCoursesOp); err != nil {
		panic(err)
	}

	appAddFavoriteCourseOp, err := reflector.NewOperationContext(http.MethodPost, "https://ggeco-func.azurewebsites.net/app/course/favorite")
	if err != nil {
		panic(err)
	}
	appAddFavoriteCourseOp.SetDescription("Add Favorite Course")
	appAddFavoriteCourseOp.SetSummary("Add Favorite Course")
	appAddFavoriteCourseOp.AddReqStructure(app.AddFavoriteCourseRequest{})
	appAddFavoriteCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusOK
		cu.Description = "Add Favorite Course Success."
	})
	appAddFavoriteCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appAddFavoriteCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appAddFavoriteCourseOp); err != nil {
		panic(err)
	}

	appDeleteFavoriteCourseOp, err := reflector.NewOperationContext(http.MethodDelete, "https://ggeco-func.azurewebsites.net/app/course/favorite")
	if err != nil {
		panic(err)
	}
	appDeleteFavoriteCourseOp.SetDescription("Delete Favorite Course")
	appDeleteFavoriteCourseOp.SetSummary("Delete Favorite Course")
	appDeleteFavoriteCourseOp.AddReqStructure(app.RemoveFavoriteCourseRequest{})
	appDeleteFavoriteCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusOK
		cu.Description = "Delete Favorite Course Success."
	})
	appDeleteFavoriteCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appDeleteFavoriteCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appDeleteFavoriteCourseOp); err != nil {
		panic(err)
	}

	appCheckFavoriteCourseOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/course/favorite/check")
	if err != nil {
		panic(err)
	}
	appCheckFavoriteCourseOp.SetDescription("Check Favorite Course")
	appCheckFavoriteCourseOp.SetSummary("Check Favorite Course")
	appCheckFavoriteCourseOp.AddReqStructure(app.IsFavoriteCourseRequest{})
	appCheckFavoriteCourseOp.AddRespStructure(app.IsFavoriteCourseResponse{})
	appCheckFavoriteCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appCheckFavoriteCourseOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appCheckFavoriteCourseOp); err != nil {
		panic(err)
	}

	appSearchPlacesOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/place/search")
	if err != nil {
		panic(err)
	}
	appSearchPlacesOp.SetDescription("Search Places")
	appSearchPlacesOp.SetSummary("Search Places")
	appSearchPlacesOp.AddReqStructure(app.SearchPlacesRequest{})
	appSearchPlacesOp.AddRespStructure(app.SearchPlacesResponse{})
	appSearchPlacesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appSearchPlacesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appSearchPlacesOp); err != nil {
		panic(err)
	}

	appGetPlaceOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/place")
	if err != nil {
		panic(err)
	}
	appGetPlaceOp.SetDescription("Get Place Info")
	appGetPlaceOp.SetSummary("Get Place Info")
	appGetPlaceOp.AddReqStructure(app.GetPlaceInfoRequest{})
	appGetPlaceOp.AddRespStructure(app.GetPlaceInfoResponse{})
	appGetPlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetPlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetPlaceOp); err != nil {
		panic(err)
	}

	appGetFavoritePlacesOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/place/favorite")
	if err != nil {
		panic(err)
	}
	appGetFavoritePlacesOp.SetDescription("Get Favorite Places")
	appGetFavoritePlacesOp.SetSummary("Get Favorite Places")
	appGetFavoritePlacesOp.AddRespStructure(app.GetFavoritePlacesResponse{})
	appGetFavoritePlacesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appGetFavoritePlacesOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appGetFavoritePlacesOp); err != nil {
		panic(err)
	}

	appAddFavoritePlaceOp, err := reflector.NewOperationContext(http.MethodPost, "https://ggeco-func.azurewebsites.net/app/place/favorite")
	if err != nil {
		panic(err)
	}
	appAddFavoritePlaceOp.SetDescription("Add Favorite Place")
	appAddFavoritePlaceOp.SetSummary("Add Favorite Place")
	appAddFavoritePlaceOp.AddReqStructure(app.AddFavoritePlaceRequest{})
	appAddFavoritePlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusOK
		cu.Description = "Add Favorite Place Success."
	})
	appAddFavoritePlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appAddFavoritePlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appAddFavoritePlaceOp); err != nil {
		panic(err)
	}

	appDeleteFavoritePlaceOp, err := reflector.NewOperationContext(http.MethodDelete, "https://ggeco-func.azurewebsites.net/app/place/favorite")
	if err != nil {
		panic(err)
	}
	appDeleteFavoritePlaceOp.SetDescription("Delete Favorite Place")
	appDeleteFavoritePlaceOp.SetSummary("Delete Favorite Place")
	appDeleteFavoritePlaceOp.AddReqStructure(app.RemoveFavoritePlaceRequest{})
	appDeleteFavoritePlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusOK
		cu.Description = "Delete Favorite Place Success."
	})
	appDeleteFavoritePlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appDeleteFavoritePlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appDeleteFavoritePlaceOp); err != nil {
		panic(err)
	}

	appCheckFavoritePlaceOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/place/favorite/check")
	if err != nil {
		panic(err)
	}
	appCheckFavoritePlaceOp.SetDescription("Check Favorite Place")
	appCheckFavoritePlaceOp.SetSummary("Check Favorite Place")
	appCheckFavoritePlaceOp.AddReqStructure(app.IsFavoritePlaceRequest{})
	appCheckFavoritePlaceOp.AddRespStructure(app.IsFavoritePlaceResponse{})
	appCheckFavoritePlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	appCheckFavoritePlaceOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(appCheckFavoritePlaceOp); err != nil {
		panic(err)
	}

	getUserOp, err := reflector.NewOperationContext(http.MethodGet, "https://ggeco-func.azurewebsites.net/app/profile")
	if err != nil {
		panic(err)
	}
	getUserOp.SetDescription("Get User Info")
	getUserOp.SetSummary("Get User Info")
	getUserOp.AddRespStructure(app.GetProfileResponse{})
	getUserOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	getUserOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(getUserOp); err != nil {
		panic(err)
	}

	updateUserNicknameOp, err := reflector.NewOperationContext(http.MethodPost, "https://ggeco-func.azurewebsites.net/app/profile/nickname")
	if err != nil {
		panic(err)
	}
	updateUserNicknameOp.SetDescription("Update User Nickname")
	updateUserNicknameOp.SetSummary("Update User Nickname")
	updateUserNicknameOp.AddReqStructure(app.UpdateNicknameRequest{})
	updateUserNicknameOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusOK
		cu.Description = "Update User Nickname Success."
	})
	updateUserNicknameOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	updateUserNicknameOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(updateUserNicknameOp); err != nil {
		panic(err)
	}

	updateUserBadgeOp, err := reflector.NewOperationContext(http.MethodPost, "https://ggeco-func.azurewebsites.net/app/profile/badge")
	if err != nil {
		panic(err)
	}
	updateUserBadgeOp.SetDescription("Update User Badge")
	updateUserBadgeOp.SetSummary("Update User Badge")
	updateUserBadgeOp.AddReqStructure(app.UpdateBadgeRequest{})
	updateUserBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusOK
		cu.Description = "Update User Badge Success."
	})
	updateUserBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error with Error Message."
		cu.ContentType = "text/plain"
	})
	updateUserBadgeOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request with Error Message."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(updateUserBadgeOp); err != nil {
		panic(err)
	}

	uploadImageOp, err := reflector.NewOperationContext(http.MethodPut, "https://ggeco-image-entry.azurewebsites.net/api/upload")
	if err != nil {
		panic(err)
	}
	uploadImageOp.SetDescription("Upload Image")
	uploadImageOp.SetSummary("Upload Image")
	uploadImageOp.AddReqStructure(app.UploadImageRequest{})
	uploadImageOp.AddRespStructure(app.UploadImageResponse{})
	uploadImageOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusInternalServerError
		cu.Description = "Internal Server Error."
		cu.ContentType = "text/plain"
	})
	uploadImageOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusUnauthorized
		cu.Description = "Unauthorized."
		cu.ContentType = "text/plain"
	})
	uploadImageOp.AddRespStructure(nil, func(cu *openapi.ContentUnit) {
		cu.HTTPStatus = http.StatusBadRequest
		cu.Description = "Bad Request."
		cu.ContentType = "text/plain"
	})
	if err := reflector.AddOperation(uploadImageOp); err != nil {
		panic(err)
	}

	value, err := reflector.Spec.MarshalYAML()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./function/swagger/swagger.yaml")
	if err != nil {
		panic(err)
	}

	_, err = f.Write(value)
	if err != nil {
		panic(err)
	}
}
