package courses

import "github.com/snowmerak/ggeco/lib/client/sqlserver"

type GetCourseRequest struct {
	CourseID   *string `query:"course_id"`
	AuthorID   *string `query:"author_id"`
	CourseName *string `query:"course_name"`
}

func Get(container sqlserver.Container, id sqlserver.UUID) (result Course, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [author_id], [name], [reg_date], [review] from [dbo].[Courses] WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.Id, &result.AuthorID, &result.Name, &result.RegDate, &result.Review); err != nil {
		return result, err
	}

	return
}

func GetByAuthor(container sqlserver.Container, author sqlserver.UUID) (result []Course, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [author_id], [name], [reg_date], [review] from [dbo].[Courses] WHERE [author_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(author)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rs := Course{}
		if err := rows.Scan(&rs.Id, &rs.AuthorID, &rs.Name, &rs.RegDate, &rs.Review); err != nil {
			return result, err
		}
		result = append(result, rs)
	}

	return
}

func GetByAuthorAndName(container sqlserver.Container, author sqlserver.UUID, name string) (result []Course, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [author_id], [name], [reg_date], [review] from [dbo].[Courses] WHERE [author_id] = @P1 AND [name] = @P2")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(author, name)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rs := Course{}
		if err := rows.Scan(&rs.Id, &rs.AuthorID, &rs.Name, &rs.RegDate, &rs.Review); err != nil {
			return result, err
		}
		result = append(result, rs)
	}

	return
}
