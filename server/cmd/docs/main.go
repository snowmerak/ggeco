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

	appSigninOp, err := reflector.NewOperationContext(http.MethodPost, "/app/auth/signin")
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

	appRefreshOp, err := reflector.NewOperationContext(http.MethodPost, "/app/auth/refresh")
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
