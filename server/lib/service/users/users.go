package users

import (
	"database/sql"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"log"
	"time"
)

type User struct {
	Id         sqlserver.UUID `json:"id,omitempty"`
	Nickname   string         `json:"nickname,omitempty"`
	Age        *uint8         `json:"age,omitempty"`
	Gender     *uint8         `json:"gender,omitempty"`
	CreateAt   string         `json:"create_at,omitempty"`
	LastSignin string         `json:"last_signin,omitempty"`
}

const createTable = `
CREATE TABLE [dbo].[Users] (
	[id]          UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_Users_id] DEFAULT (newid()) NOT NULL,
	[nickname]    NCHAR (18)       NOT NULL,
	[age]         TINYINT          NULL,
	[gender]      TINYINT          NULL,
	[create_at]   DATE             NOT NULL,
	[last_signin] DATE             NOT NULL,
	CONSTRAINT [PK_Users] PRIMARY KEY CLUSTERED ([id] ASC)
);
`

func CreateTable(container sqlserver.Container) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	_, err = client.Exec(createTable)

	return
}

func Get(container sqlserver.Container, id string) (result User, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	row := client.QueryRow("SELECT [id], [nickname], [age], [gender] from [dbo].[Users]")
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.Id, &result.Nickname, &result.Age, &result.Gender); err != nil {
		return result, err
	}

	return
}

const insertUser = `DECLARE @insertedId TABLE (id UNIQUEIDENTIFIER)
INSERT INTO [dbo].[Users] ([nickname], [create_at], [last_signin]) OUTPUT INSERTED.id VALUES (@P1, @P2, @P2)
SELECT id FROM @insertedId`

func Add(container sqlserver.Container, nickname string) (id sqlserver.UUID, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(insertUser)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Println(err)
		}
	}(stmt)

	t := time.Now().Format(sqlserver.DateFormat)
	row := stmt.QueryRow(nickname, t)
	if err = row.Err(); err != nil {
		return
	}

	if err = row.Scan(&id); err != nil {
		return
	}

	return
}

const updateAge = `UPDATE [dbo].[Users] SET [age] = @P1 WHERE [id] = @P2`

func UpdateAge(container sqlserver.Container, id sqlserver.UUID, age uint8) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(updateAge)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(age, id)
	if err != nil {
		return
	}

	return
}

const updateGender = `UPDATE [dbo].[Users] SET [gender] = @P1 WHERE [id] = @P2`

func UpdateGender(container sqlserver.Container, id sqlserver.UUID, gender uint8) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(updateGender)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(gender, id)
	if err != nil {
		return
	}

	return
}

const updateLastSignin = `UPDATE [dbo].[Users] SET [last_signin] = @P1 WHERE [id] = @P2`

func UpdateLastSignin(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(updateLastSignin)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(time.Now().Format(sqlserver.DateFormat), id)

	return
}

const deleteUser = `DELETE FROM [dbo].[Users] WHERE [id] = @P1`

func Delete(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(deleteUser)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(id)

	return
}

const updateNickname = `UPDATE [dbo].[Users] SET [nickname] = @P1 WHERE [id] = @P2`

func UpdateNickname(container sqlserver.Container, id sqlserver.UUID, nickname string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(updateNickname)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(nickname, id)

	return
}
