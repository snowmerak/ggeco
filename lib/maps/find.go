package maps

import (
	"context"
	"googlemaps.github.io/maps"
)

// SearchPlaceIdRequest is a request for SearchPlaceId
// jetti:parameter
type SearchPlaceIdRequest struct {
	PlaceID  string
	Language string

	GetReviews           bool
	GetOpeningHours      bool
	GetPermanentlyClosed bool
	GetGeometryLocation  bool
	GetPhotos            bool
	GetPhone             bool
}

// SearchPlaceIdResponse is a response for SearchPlaceId
type SearchPlaceIdResponse struct {
	PlaceID          string       `json:"place_id,omitempty"`
	FormattedAddress string       `json:"formatted_address,omitempty"`
	GeometryLocation Location     `json:"geometry_location,omitempty"`
	Name             string       `json:"name,omitempty"`
	OpeningHours     []TimeShift  `json:"opening_hours,omitempty"`
	BusinessStatus   string       `json:"business_status,omitempty"`
	Reviews          []Review     `json:"reviews,omitempty"`
	Photos           []PlacePhoto `json:"photos,omitempty"`
	Phone            string       `json:"phone,omitempty"`
}

type Location struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

type TimeShift struct {
	OpenDay   string `json:"open_day,omitempty"`
	OpenTime  string `json:"open_time,omitempty"`
	CloseDay  string `json:"close_day,omitempty"`
	CloseTime string `json:"close_time,omitempty"`
}

type Review struct {
	AuthorName string `json:"author_name,omitempty"`
	Language   string `json:"language,omitempty"`
	Rating     int    `json:"rating,omitempty"`
	Text       string `json:"text,omitempty"`
}

type PlacePhoto struct {
	PhotoReference   string   `json:"photo_reference,omitempty"`
	Height           int      `json:"height,omitempty"`
	Width            int      `json:"width,omitempty"`
	HTMLAttributions []string `json:"html_attributions,omitempty"`
	PhotoPath        string   `json:"photo_path,omitempty"`
}

func SearchPlaceId(ctx context.Context, container Container, requestOpt ...SearchPlaceIdRequestOptional) (response SearchPlaceIdResponse, err error) {
	client, err := GetClient(container)
	if err != nil {
		return response, err
	}

	request := ApplySearchPlaceIdRequest(SearchPlaceIdRequest{}, requestOpt...)
	if request.Language == "" {
		request.Language = "ko"
	}

	fields := []maps.PlaceDetailsFieldMask{
		maps.PlaceDetailsFieldMaskFormattedAddress,
		maps.PlaceDetailsFieldMaskName,
	}
	if request.GetReviews {
		fields = append(fields, maps.PlaceDetailsFieldMaskReviews)
	}
	if request.GetOpeningHours {
		fields = append(fields, maps.PlaceDetailsFieldMaskOpeningHours)
	}
	if request.GetPermanentlyClosed {
		fields = append(fields, maps.PlaceDetailsFieldMaskBusinessStatus)
	}
	if request.GetGeometryLocation {
		fields = append(fields, maps.PlaceDetailsFieldMaskGeometryLocation)
	}
	if request.GetPhotos {
		fields = append(fields, maps.PlaceDetailsFieldMaskPhotos)
	}
	if request.GetPhone {
		fields = append(fields, maps.PlaceDetailsFieldMaskFormattedPhoneNumber)
	}

	result, err := client.baseClient.PlaceDetails(ctx, &maps.PlaceDetailsRequest{
		PlaceID:  request.PlaceID,
		Language: request.Language,
		Fields:   fields,
	})
	if err != nil {
		return response, err
	}

	response.Name = result.Name
	response.FormattedAddress = result.FormattedAddress
	response.GeometryLocation = Location{
		Lat: result.Geometry.Location.Lat,
		Lng: result.Geometry.Location.Lng,
	}
	response.BusinessStatus = result.BusinessStatus
	for _, review := range result.Reviews {
		response.Reviews = append(response.Reviews, Review{
			AuthorName: review.AuthorName,
			Language:   review.Language,
			Rating:     review.Rating,
			Text:       review.Text,
		})
	}

	if result.OpeningHours != nil {
		for _, period := range result.OpeningHours.Periods {
			response.OpeningHours = append(response.OpeningHours, TimeShift{
				OpenDay:   period.Open.Day.String(),
				OpenTime:  period.Open.Time,
				CloseDay:  period.Close.Day.String(),
				CloseTime: period.Close.Time,
			})
		}
	}

	if result.Photos != nil {
		for _, photo := range result.Photos {
			response.Photos = append(response.Photos, PlacePhoto{
				PhotoReference:   photo.PhotoReference,
				Height:           photo.Height,
				Width:            photo.Width,
				HTMLAttributions: photo.HTMLAttributions,
				PhotoPath:        client.SignPhotoURL(photo.PhotoReference),
			})
		}
	}

	if result.FormattedPhoneNumber != "" {
		response.Phone = result.FormattedPhoneNumber
	}

	return response, nil
}
