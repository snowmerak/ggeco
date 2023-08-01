package courses

import "github.com/snowmerak/ggeco/lib/client/sqlserver"

func Add(container sqlserver.Container, author sqlserver.UUID, name string, regDate string, review string) (id sqlserver.UUID, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	row := client.QueryRow(`DECLARE @InsertedId TABLE (id uniqueidentifier)
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
