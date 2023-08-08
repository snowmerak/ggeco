package users

import "github.com/snowmerak/ggeco/server/lib/client/sqlserver"

type Kakao struct {
	UserId  sqlserver.UUID `json:"user_id,omitempty"`
	KakaoId int64          `json:"kakao_id,omitempty"`
	Info    string         `json:"info,omitempty"`
}

type GetKakaoUserRequest struct {
	Id string `query:"id" required:"true"`
}

type GetKakaoUserResponse struct {
	Id      string `json:"id"`
	KakaoId int64  `json:"kakao_id"`
	Info    string `json:"info"`
}

func GetKakaoUser(container sqlserver.Container, id sqlserver.UUID) (result Kakao, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	result.UserId = id

	stmt, err := client.Prepare("SELECT [kakao_id], [info] from [dbo].[KakaoUsers] WHERE [user_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.KakaoId, &result.Info); err != nil {
		return result, err
	}

	result.UserId = id

	return
}

type GetKakaoUserByKakaoIdRequest struct {
	KakaoId int64 `query:"kakao_id" required:"true"`
}

type GetKakaoUserByKakaoIdResponse struct {
	Id      string `json:"id"`
	KakaoId int64  `json:"kakao_id"`
	Info    string `json:"info"`
}

func GetKakaoUserByKakaoId(container sqlserver.Container, kakaoId int64) (result Kakao, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [user_id], [info] from [dbo].[KakaoUsers] WHERE [kakao_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(kakaoId)
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.UserId, &result.Info); err != nil {
		return result, err
	}

	result.KakaoId = kakaoId

	return
}

type AddKakaoUserRequest struct {
	UserId  sqlserver.UUID `json:"user_id" required:"true"`
	KakaoId int64          `json:"kakao_id" required:"true"`
	Info    string         `json:"info" required:"true"`
}

func AddKakaoUser(container sqlserver.Container, userId sqlserver.UUID, kakaoId int64, info string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("INSERT INTO [dbo].[KakaoUsers] ([user_id], [kakao_id], [info]) VALUES (@P1, @P2, @P3)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, kakaoId, info)

	return
}

type UpdateKakaoUserRequest struct {
	Id      sqlserver.UUID `json:"id" required:"true"`
	KakaoId int64          `json:"kakao_id" required:"true"`
	Info    string         `json:"info" required:"true"`
}

func UpdateKakaoUser(container sqlserver.Container, id sqlserver.UUID, kakaoId int64, info string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[KakaoUsers] SET [kakao_id] = @P1, [info] = @P2 WHERE [user_id] = @P3")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(kakaoId, info, id)

	return
}

type DeleteKakaoUserRequest struct {
	Id sqlserver.UUID `query:"id" required:"true"`
}

func DeleteKakaoUser(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("DELETE FROM [dbo].[KakaoUsers] WHERE [user_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return
}
