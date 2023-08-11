package courses

import (
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"github.com/snowmerak/ggeco/server/lib/service/place"
)

type Place struct {
	Id       sqlserver.UUID `json:"id"`
	CourseId sqlserver.UUID `json:"course_id"`
	PlaceId  string         `json:"place_id"`
	Order    int            `json:"order"`
}

type GetPlacesRequest struct {
	CourseId string `query:"course_id" required:"true"`
}

type GetPlacesResponse struct {
	Places []place.GetPlaceResponse `json:"places"`
}

func GetPlaces(container sqlserver.Container, courseId sqlserver.UUID) (result []Place, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [course_id], [place_id], [order] from [dbo].[CoursePlaces] WHERE [course_id] = @P1 ORDER BY [order]")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(courseId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Place
		err = rows.Scan(&p.Id, &p.CourseId, &p.PlaceId, &p.Order)
		if err != nil {
			return
		}
		result = append(result, p)
	}

	return
}

type SetPlacesRequest struct {
	CourseId string   `query:"course_id" required:"true"`
	PlaceIds []string `query:"place_ids" required:"true"`
}

func SetPlaces(container sqlserver.Container, courseId sqlserver.UUID, places []string) error {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	tx, err := client.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM [dbo].[CoursePlaces] WHERE [course_id] = @P1", courseId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO [dbo].[CoursePlaces] ([course_id], [place_id], [order]) VALUES (@P1, @P2, @P3)")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	for i, placeId := range places {
		_, err = stmt.Exec(courseId, placeId, i)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}
	}

	return tx.Commit()
}

func GetCoursesFromPlace(container sqlserver.Container, placeId string, count int) ([]Course, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return nil, err
	}

	stmt, err := client.Prepare(`SELECT TOP (@P1) [course_id] FROM [dbo].[CoursePlaces] WHERE [place_id] = @P2 ORDER BY RAND()`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(count, placeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Course, 0, count)
	for rows.Next() {
		courseId := sqlserver.UUID{}
		if err := rows.Scan(&courseId); err != nil {
			return result, err
		}

		rs, err := Get(container, courseId)
		if err != nil {
			return result, err
		}

		result = append(result, rs)
	}

	return result, nil
}
