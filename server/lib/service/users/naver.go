package users

import "github.com/snowmerak/ggeco/server/lib/client/sqlserver"

type Naver struct {
	Id      sqlserver.UUID `json:"id,omitempty"`
	NaverId string         `json:"naver_id,omitempty"`
	Info    string         `json:"info,omitempty"`
}

const createNaverTable = `
CREATE TABLE [dbo].[NaverUsers] (
    [user_id]  UNIQUEIDENTIFIER NOT NULL,
    [naver_id] CHAR (128)       NOT NULL,
    [info]     NTEXT            NOT NULL,
    CONSTRAINT [PK_NaverUsers] PRIMARY KEY CLUSTERED ([user_id] ASC)
);


GO
CREATE NONCLUSTERED INDEX [index_naver_id]
    ON [dbo].[NaverUsers]([naver_id] ASC);
`

func CreateNaverTable(container sqlserver.Container) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	_, err = client.Exec(createNaverTable)

	return
}

func GetNaver(container sqlserver.Container, id sqlserver.UUID) (result Naver, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	result.Id = id

	stmt, err := client.Prepare("SELECT [naver_id], [info] from [dbo].[NaverUsers] WHERE [user_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.NaverId, &result.Info); err != nil {
		return result, err
	}

	return
}

func GetNaverByNaverId(container sqlserver.Container, naverId string) (result Naver, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [user_id], [info] from [dbo].[NaverUsers] WHERE [naver_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(naverId)
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.Id, &result.Info); err != nil {
		return result, err
	}

	return
}

func AddNaver(container sqlserver.Container, id sqlserver.UUID, naverId string, info string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`INSERT INTO [dbo].[NaverUsers] ([user_id], [naver_id], [info]) VALUES (@P1, @P2, @P3)`)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, naverId, info)

	return
}

func UpdateNaver(container sqlserver.Container, id sqlserver.UUID, naverId string, info string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`UPDATE [dbo].[NaverUsers] SET [naver_id] = @P2, [info] = @P3 WHERE [user_id] = @P1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, naverId, info)

	return
}

func DeleteNaver(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`DELETE FROM [dbo].[NaverUsers] WHERE [user_id] = @P1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return
}
