package users

import "github.com/snowmerak/ggeco/server/lib/client/sqlserver"

type Naver struct {
	Id      sqlserver.UUID `json:"id,omitempty"`
	NaverId string         `json:"naver_id,omitempty"`
	Info    string         `json:"info,omitempty"`
}

type GetNaverUserRequest struct {
	Id string `query:"id" required:"true"`
}

type GetNaverUserResponse struct {
	Id      string `json:"id"`
	NaverId string `json:"naver_id"`
	Info    string `json:"info"`
}

func GetNaverUser(container sqlserver.Container, id sqlserver.UUID) (result Naver, err error) {
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

type GetNaverUserByNaverIdRequest struct {
	NaverId string `query:"naver_id" required:"true"`
}

type GetNaverUserByNaverIdResponse struct {
	Id      string `json:"id"`
	NaverId string `json:"naver_id"`
	Info    string `json:"info"`
}

func GetNaverUserByNaverId(container sqlserver.Container, naverId string) (result Naver, err error) {
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

	result.NaverId = naverId

	return
}

type AddNaverUserRequest struct {
	Id      string `json:"id"`
	NaverId string `json:"naver_id"`
	Info    string `json:"info"`
}

func AddNaverUser(container sqlserver.Container, id sqlserver.UUID, naverId string, info string) (err error) {
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

type UpdateNaverUserRequest struct {
	Id      string `json:"id"`
	NaverId string `json:"naver_id"`
	Info    string `json:"info"`
}

func UpdateNaverUser(container sqlserver.Container, id sqlserver.UUID, naverId string, info string) (err error) {
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

type DeleteNaverUserRequest struct {
	Id string `query:"id" required:"true"`
}

func DeleteNaverUser(container sqlserver.Container, id sqlserver.UUID) (err error) {
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
