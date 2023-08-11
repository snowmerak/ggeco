package place

import (
	"context"
	"github.com/snowmerak/ggeco/server/lib/client/maps"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
)

type SearchTextRequest struct {
	Query     string  `query:"query" required:"true"`
	Lang      string  `query:"lang"`
	Latitude  float64 `query:"latitude"`
	Longitude float64 `query:"longitude"`
	Radius    int64   `query:"radius"`
	OpenNow   bool    `query:"opennow"`
}

func SearchText(ctx context.Context, container maps.Container, query string, radius int64, lang string, latitude float64, longitude float64, openNow bool) (response []*maps.SearchTextResponse, err error) {
	resp, err := maps.SearchText(ctx, container, func(request *maps.SearchTextRequest) *maps.SearchTextRequest {
		request.Query = query
		request.Radius = radius
		request.Language = lang
		request.Latitude = latitude
		request.Longitude = longitude
		request.OpenNow = openNow
		return request
	})
	if err != nil {
		return response, err
	}

	return resp, nil
}

func GetCourseContainsPlace(container maps.Container, placeId string, count int) ([]sqlserver.UUID, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return nil, err
	}

	stmt, err := client.Prepare("SELECT TOP (@P2) [course_id] FROM [dbo].[CoursePlaces] WHERE [place_id] = @P1 ORDER BY RAND()")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(placeId, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	step := 0
	var result []sqlserver.UUID
	for rows.Next() && step < count {
		var courseId sqlserver.UUID
		err = rows.Scan(&courseId)
		if err != nil {
			return nil, err
		}
		result = append(result, courseId)
		step++
	}

	return result, nil
}
