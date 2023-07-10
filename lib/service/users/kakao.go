package users

import "github.com/snowmerak/ggeco/lib/client/sqlserver"

type Kakao struct {
	UserId  sqlserver.UUID `json:"user_id,omitempty"`
	KakaoId int64          `json:"kakao_id,omitempty"`
	Info    string         `json:"info,omitempty"`
}

const createKakaoTable = `
CREATE TABLE [dbo].[KakaoUsers] (
    [user_id]  UNIQUEIDENTIFIER NOT NULL,
    [kakao_id] BIGINT           NOT NULL,
    [info]     NTEXT            NOT NULL,
    CONSTRAINT [PK_KakaoUsers] PRIMARY KEY CLUSTERED ([user_id] ASC)
);


GO
CREATE NONCLUSTERED INDEX [index_kakao_id]
    ON [dbo].[KakaoUsers]([kakao_id] ASC);
`

func CreateKakaoTable(container sqlserver.Container) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	_, err = client.Exec(createKakaoTable)

	return
}

func GetKakao(container sqlserver.Container, id sqlserver.UUID) (result Kakao, err error) {
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

	return
}

func GetKakaoByKakaoId(container sqlserver.Container, kakaoId int64) (result Kakao, err error) {
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

	return
}

func AddKakao(container sqlserver.Container, kakaoId int64, info string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("INSERT INTO [dbo].[KakaoUsers] ([user_id], [kakao_id], [info]) VALUES (NEWID(), @P1, @P2)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(kakaoId, info)

	return
}

func UpdateKakao(container sqlserver.Container, id sqlserver.UUID, kakaoId int64, info string) (err error) {
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

func DeleteKakao(container sqlserver.Container, id sqlserver.UUID) (err error) {
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
