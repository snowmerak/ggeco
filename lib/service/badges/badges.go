package badges

import "github.com/snowmerak/ggeco/lib/client/sqlserver"

type Badge struct {
	Id      sqlserver.UUID `json:"id,omitempty"`
	Name    string         `json:"name,omitempty"`
	Summary string         `json:"summary,omitempty"`
}

const createTable = `
CREATE TABLE [dbo].[Badges] (
	[id]      UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_Badges_id] DEFAULT (newid()) NOT NULL,
	[name]    NCHAR (12)       NOT NULL,
	[summary] NTEXT            NOT NULL,
	CONSTRAINT [PK_Badges] PRIMARY KEY CLUSTERED ([id] ASC)
);

GO
CREATE NONCLUSTERED INDEX [Index_name]
    ON [dbo].[Badges]([name] ASC);
`

func CreateTable(container sqlserver.Container) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	_, err = client.Exec(createTable)

	return
}

func Get(container sqlserver.Container, id string) (result Badge, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	row := client.QueryRow("SELECT [id], [name], [summary] from [dbo].[Badges]")
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.Id, &result.Name, &result.Summary); err != nil {
		return result, err
	}

	return
}

func Add(container sqlserver.Container, name string, summary string) (id sqlserver.UUID, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`DECLARE @InsertedId TABLE (id UNIQUEIDENTIFIER)
INSERT INTO [dbo].[Badges] ([name], [summary]) OUTPUT INSERTED.[id] INTO @InsertedId VALUES (@P1, @P2)
SELECT id from @InsertedId
`)
	if err != nil {
		return
	}

	row := stmt.QueryRow(name, summary)
	if err := row.Err(); err != nil {
		return id, err
	}

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return
}

func UpdateName(container sqlserver.Container, id string, name string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Badges] SET [name] = @P2 WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name)

	return
}

func UpdateSummary(container sqlserver.Container, id string, summary string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Badges] SET [summary] = @P2 WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, summary)

	return
}

func Delete(container sqlserver.Container, id string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("DELETE FROM [dbo].[Badges] WHERE [id] = @P1")
	if err != nil {
		return
	}

	_, err = stmt.Exec(id)

	return
}
