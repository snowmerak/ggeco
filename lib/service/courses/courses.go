package courses

import "github.com/snowmerak/ggeco/lib/client/sqlserver"

type Course struct {
	Id      sqlserver.UUID `json:"id,omitempty"`
	Author  sqlserver.UUID `json:"author,omitempty"`
	Name    string         `json:"name,omitempty"`
	RegDate string         `json:"reg_date,omitempty"`
	Review  string         `json:"review,omitempty"`
}

const createTable = `
CREATE TABLE [dbo].[Courses] (
    [id]       UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_Courses_id] DEFAULT (newid()) NOT NULL,
    [author]   UNIQUEIDENTIFIER NOT NULL,
    [name]     NCHAR (40)       NOT NULL,
    [reg_date] DATE             NOT NULL,
    [review]   NTEXT            NOT NULL,
    CONSTRAINT [PK_Courses] PRIMARY KEY CLUSTERED ([id] ASC)
);


GO
CREATE NONCLUSTERED INDEX [index_author]
    ON [dbo].[Courses]([author] ASC);


GO
CREATE NONCLUSTERED INDEX [Index_author_date]
    ON [dbo].[Courses]([author] ASC, [reg_date] ASC);
`

func CreateTable(container sqlserver.Container) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	_, err = client.Exec(createTable)

	return
}

func Get(container sqlserver.Container, id string) (result Course, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	row := client.QueryRow("SELECT [id], [author], [name], [reg_date], [review] from [dbo].[Courses]")
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.Id, &result.Author, &result.Name, &result.RegDate, &result.Review); err != nil {
		return result, err
	}

	return
}

func Add(container sqlserver.Container, author sqlserver.UUID, name string, regDate string, review string) (id sqlserver.UUID, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	row := client.QueryRow(`DECLARE @InsertedId TABLE (id uniqueidentifier)
INSERT INTO [dbo].[Courses] ([author], [name], [reg_date], [review]) OUTPUT INSERTED.id VALUES (@P1, @P2, @P3, @P4)
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

func Delete(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("DELETE FROM [dbo].[Courses] WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return
}
