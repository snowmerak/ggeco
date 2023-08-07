package place

import "github.com/snowmerak/ggeco/server/lib/client/sqlserver"

type Review struct {
	Id        sqlserver.UUID `json:"id,omitempty"`
	PlaceId   string         `json:"place_id,omitempty"`
	CourseId  sqlserver.UUID `json:"course_id,omitempty"`
	AuthorId  sqlserver.UUID `json:"author_id,omitempty"`
	Latitude  float64        `json:"latitude,omitempty"`
	Longitude float64        `json:"longitude,omitempty"`
	Review    string         `json:"review,omitempty"`
}

type GetReviewRequest struct {
	CourseId string `query:"course_id" required:"true"`
}

type GetReviewResponse struct {
	Id        string  `json:"id,omitempty"`
	AuthorId  string  `json:"author_id,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Review    string  `json:"review,omitempty"`
}

func GetReviewsOfCourse(container sqlserver.Container, courseId sqlserver.UUID) (result []Review, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [course_id], [place_id], [author_id], [latitude], [longitude], [review] from [dbo].[PlaceReviews] WHERE [course_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(courseId)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Review
		err = rows.Scan(&r.Id, &r.CourseId, &r.PlaceId, &r.AuthorId, &r.Latitude, &r.Longitude, &r.Review)
		if err != nil {
			return
		}
		result = append(result, r)
	}

	return
}

type UpdateReviewRequest struct {
	Id     string `json:"id"`
	Review string `json:"review"`
}

func UpdateReview(container sqlserver.Container, id sqlserver.UUID, review string) error {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	stmt, err := client.Prepare("UPDATE [dbo].[PlaceReviews] SET [review] = @P1 WHERE [course_id] = @P2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(review, id)

	return err
}

type DeleteReviewRequest struct {
	Id string `query:"id" required:"true"`
}

func DeleteReview(container sqlserver.Container, id sqlserver.UUID) error {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	stmt, err := client.Prepare("DELETE FROM [dbo].[PlaceReviews] WHERE [id] = @P1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

type CreateReviewRequest struct {
	CourseId  string  `json:"course_id"`
	PlaceId   string  `json:"place_id"`
	AuthorId  string  `json:"author_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Review    string  `json:"review"`
}

type CreateReviewResponse struct {
	Id string `json:"id"`
}

func CreateReview(container sqlserver.Container, courseId sqlserver.UUID, placeId string, authorId sqlserver.UUID, latitude float64, longitude float64, review string) (sqlserver.UUID, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return sqlserver.UUID{}, err
	}

	stmt, err := client.Prepare(`DECLARE @insertedId uniqueidentifier
INSERT INTO [dbo].[PlaceReviews] ([course_id], [place_id], [author_id], [latitude], [longitude], [review]) OUTPUT inserted.id VALUES (@P1, @P2, @P3, @P4, @P5, @P6)
SELECT @insertedId
`)
	if err != nil {
		return sqlserver.UUID{}, err
	}
	defer stmt.Close()

	rs := stmt.QueryRow(courseId, placeId, authorId, latitude, longitude, review)
	if err := rs.Err(); err != nil {
		return sqlserver.UUID{}, err
	}

	var id sqlserver.UUID
	err = rs.Scan(&id)
	if err != nil {
		return sqlserver.UUID{}, err
	}

	return id, nil
}
