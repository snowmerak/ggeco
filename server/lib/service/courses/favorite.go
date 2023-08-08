package courses

import (
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"time"
)

type FavoriteCourse struct {
	Id           sqlserver.UUID `json:"id"`
	UserId       sqlserver.UUID `json:"user_id"`
	CourseId     sqlserver.UUID `json:"course_id"`
	RegisteredAt time.Time      `json:"registered_at"`
}

type GetFavoriteCoursesByUserIdRequest struct {
	UserId string `query:"user_id"`
}

type GetFavoriteCourseByUserIdResponse struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	CourseId     string    `json:"course_id"`
	RegisteredAt time.Time `json:"registered_at"`
}

func GetFavoriteCoursesByUserId(container sqlserver.Container, userId sqlserver.UUID) ([]FavoriteCourse, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return nil, err
	}

	var favoriteCourses []FavoriteCourse

	stmt, err := client.Prepare("SELECT [id], [user_id], [course_id], [registered_at] FROM [dbo].[FavoriteCourses] WHERE [user_id] = @P1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var favoriteCourse FavoriteCourse
		err := rows.Scan(&favoriteCourse.Id, &favoriteCourse.UserId, &favoriteCourse.CourseId, &favoriteCourse.RegisteredAt)
		if err != nil {
			return nil, err
		}
		favoriteCourses = append(favoriteCourses, favoriteCourse)
	}

	return favoriteCourses, nil
}

type CountFavoriteCourseRequest struct {
	CourseId string `query:"course_id"`
}

type CountFavoriteCourseResponse struct {
	Count int `json:"count"`
}

func CountFavoriteCourse(container bean.Container, courseId sqlserver.UUID) (int, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return 0, err
	}

	stmt, err := client.Prepare("SELECT COUNT(*) FROM [dbo].[FavoriteCourses] WHERE [course_id] = @P1")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(courseId)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

type AddFavoriteCourseRequest struct {
	UserId   string `json:"user_id"`
	CourseId string `json:"course_id"`
}

type AddFavoriteCourseResponse struct {
	Id string `json:"id"`
}

func AddFavoriteCourse(container bean.Container, userId sqlserver.UUID, courseId sqlserver.UUID) (sqlserver.UUID, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return nil, err
	}

	stmt, err := client.Prepare(`DECLARE @inserted UNIQUEIDENTIFIER
INSERT INTO [dbo].[FavoriteCourses] ([user_id], [course_id], [registered_at]) OUTPUT inserted.id VALUES (@P1, @P2, @P3)
SELECT @inserted`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId, courseId, time.Now().Format(sqlserver.DateFormat))
	if err := row.Err(); err != nil {
		return nil, err
	}

	var id sqlserver.UUID
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

type DeleteFavoriteCourseRequest struct {
	Id string `query:"id"`
}

func DeleteFavoriteCourse(container bean.Container, id sqlserver.UUID) error {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	stmt, err := client.Prepare("DELETE FROM [dbo].[FavoriteCourses] WHERE [id] = @P1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
