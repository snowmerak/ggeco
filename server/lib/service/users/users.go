package users

import (
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"time"
)

type User struct {
	Id         sqlserver.UUID `json:"id,omitempty"`
	Nickname   string         `json:"nickname,omitempty"`
	Age        *uint8         `json:"age,omitempty"`
	Gender     *uint8         `json:"gender,omitempty"`
	CreateAt   time.Time      `json:"create_at,omitempty"`
	LastSignin time.Time      `json:"last_signin,omitempty"`
	Badge      sqlserver.UUID `json:"badge,omitempty"`
}

type GetUserRequest struct {
	Id string `query:"id" required:"true"`
}

type GetUserResponse struct {
	Id         string    `json:"id"`
	Nickname   string    `json:"nickname"`
	CreateAt   time.Time `json:"create_at"`
	LastSignin time.Time `json:"last_signin"`
	Badge      string    `json:"badge"`
}

func GetUser(container sqlserver.Container, id sqlserver.UUID) (result User, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [nickname], [create_at], [last_signin], [badge] FROM [dbo].[Users] WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err = row.Err(); err != nil {
		return
	}

	err = row.Scan(&result.Id, &result.Nickname, &result.CreateAt, &result.LastSignin, &result.Badge)

	return
}

const insertUser = `DECLARE @insertedId TABLE (id UNIQUEIDENTIFIER)
INSERT INTO [dbo].[Users] ([nickname], [create_at], [last_signin]) OUTPUT INSERTED.id INTO @insertedId VALUES (@P1, @P2, @P2)
SELECT id FROM @insertedId`

type AddUserRequest struct {
	Nickname string `json:"nickname" required:"true"`
}

type AddUserResponse struct {
	Id string `json:"id"`
}

func AddUser(container sqlserver.Container, nickname string) (id sqlserver.UUID, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(insertUser)
	if err != nil {
		return
	}
	defer stmt.Close()

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

type UpdateAgeRequest struct {
	Id  string `json:"id" required:"true"`
	Age uint8  `json:"age" required:"true"`
}

func UpdateAge(container sqlserver.Container, id sqlserver.UUID, age uint8) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(updateAge)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(age, id)
	if err != nil {
		return
	}

	return
}

const updateGender = `UPDATE [dbo].[Users] SET [gender] = @P1 WHERE [id] = @P2`

type UpdateGenderRequest struct {
	Id     string `json:"id" required:"true"`
	Gender uint8  `json:"gender" required:"true"`
}

func UpdateGender(container sqlserver.Container, id sqlserver.UUID, gender uint8) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(updateGender)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(gender, id)
	if err != nil {
		return
	}

	return
}

const updateLastSignin = `UPDATE [dbo].[Users] SET [last_signin] = @P1 WHERE [id] = @P2`

type UpdateLastSigninRequest struct {
	Id string `json:"id" required:"true"`
}

func UpdateLastSignin(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(updateLastSignin)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now().Format(sqlserver.DateFormat), id)

	return
}

const deleteUser = `DELETE FROM [dbo].[Users] WHERE [id] = @P1`

type DeleteUserRequest struct {
	Id string `query:"id" required:"true"`
}

func DeleteUser(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(deleteUser)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return
}

const updateNickname = `UPDATE [dbo].[Users] SET [nickname] = @P1 WHERE [id] = @P2`

type UpdateNicknameRequest struct {
	Id       string `json:"id" required:"true"`
	Nickname string `json:"nickname" required:"true"`
}

func UpdateNickname(container sqlserver.Container, id sqlserver.UUID, nickname string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(updateNickname)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(nickname, id)

	return
}

func UpdateBadge(container sqlserver.Container, id sqlserver.UUID, badge sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Users] SET [badge] = @P1 WHERE [id] = @P2")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(badge, id)

	return
}
