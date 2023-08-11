package courses

import "github.com/snowmerak/ggeco/server/lib/client/sqlserver"

type AddCourseRequest struct {
	AuthorID string `json:"author_id" required:"true"`
	Name     string `json:"name" required:"true"`
	RegDate  string `json:"reg_date" required:"true"`
	Review   string `json:"review" required:"true"`
}

type AddCourseResponse struct {
	Id string `json:"id" required:"true"`
}

func Add(container sqlserver.Container, author sqlserver.UUID, name string, regDate string, review string) (id sqlserver.UUID, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	row := client.QueryRow(`DECLARE @InsertedId table(id uniqueidentifier)
INSERT INTO [dbo].[Courses] ([author_id], [name], [reg_date], [review]) OUTPUT INSERTED.id VALUES (@P1, @P2, @P3, @P4)
SELECT id from @InsertedId`, author, name, regDate, review)
	if err := row.Err(); err != nil {
		return id, err
	}

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return
}

type UpdateCourseNameRequest struct {
	CourseID string `json:"course_id" required:"true"`
	Name     string `json:"name" required:"true"`
}

func UpdateName(container sqlserver.Container, id sqlserver.UUID, name string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Courses] SET [name] = @P1 WHERE [id] = @P2")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, id)

	return
}

type UpdateCourseReviewRequest struct {
	CourseID string `json:"course_id" required:"true"`
	Review   string `json:"review" required:"true"`
}

func UpdateReview(container sqlserver.Container, id sqlserver.UUID, review string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Courses] SET [review] = @P1 WHERE [id] = @P2")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(review, id)

	return
}

func UpdateDate(container sqlserver.Container, id sqlserver.UUID, date string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Courses] SET [reg_date] = @P1 WHERE [id] = @P2")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(date, id)

	return
}

func UpdateIsPublic(container sqlserver.Container, id sqlserver.UUID, isPublic bool) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Courses] SET [is_public] = @P1 WHERE [id] = @P2")
	if err != nil {
		return
	}
	defer stmt.Close()

	v := 0
	if isPublic {
		v = 1
	}
	_, err = stmt.Exec(v, id)

	return
}
