package sqlserver

func GetUserInfo(container Container, id string) (result UserInfo, err error) {
	client, err := GetClient(container)
	if err != nil {
		return
	}

	row := client.baseClient.QueryRow("SELECT [id], [nickname], [age], [gender] from [dbo].[Users]")
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.Id, &result.Nickname, &result.Age, &result.Gender); err != nil {
		return result, err
	}

	return
}

func AddNewUser(container Container, nickname string) (id UUID, err error) {
	client, err := GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.baseClient.Prepare(`DECLARE @insertedId TABLE (id UNIQUEIDENTIFIER)
INSERT INTO [dbo].[Users] ([nickname]) OUTPUT INSERTED.id VALUES (@P1)
SELECT id FROM @insertedId`)
	if err != nil {
		return
	}

	row := stmt.QueryRow(nickname)
	if err = row.Err(); err != nil {
		return
	}

	if err = row.Scan(&id); err != nil {
		return
	}

	return
}
