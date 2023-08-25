package place

import "github.com/snowmerak/ggeco/server/lib/client/sqlserver"

var _ = `
CREATE TABLE [dbo].[UserPlaceVisitCount] (
    [id]         UNIQUEIDENTIFIER NOT NULL,
    [user_id]    UNIQUEIDENTIFIER NOT NULL,
    [place_type] NVARCHAR (60)    NOT NULL,
    [count]      BIGINT           NOT NULL,
    CONSTRAINT [PK_NewTable] PRIMARY KEY CLUSTERED ([id] ASC)
);
`

func GetVisitCount(container sqlserver.Container, userId sqlserver.UUID, placeType string) (count int, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [count] FROM [dbo].[UserPlaceVisitCount] WHERE [user_id] = @P1 AND [place_type] = @P2")
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(userId, placeType).Scan(&count)
	if err != nil {
		return
	}

	return
}

func AddOrInitVisitCount(container sqlserver.Container, userId sqlserver.UUID, placeType string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT COUNT(*) FROM [dbo].[UserPlaceVisitCount] WHERE [user_id] = @P1 AND [place_type] = @P2")
	if err != nil {
		return
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(userId, placeType).Scan(&count)
	if err != nil {
		return
	}

	if count > 0 {
		stmt, err = client.Prepare("UPDATE [dbo].[UserPlaceVisitCount] SET [count] = [count] + 1 WHERE [user_id] = @P1 AND [place_type] = @P2")
		if err != nil {
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(userId, placeType)
		if err != nil {
			return
		}
	} else {
		stmt, err = client.Prepare("INSERT INTO [dbo].[UserPlaceVisitCount] ([id], [user_id], [place_type], [count]) VALUES (NEWID(), @P1, @P2, 1)")
		if err != nil {
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(userId, placeType)
		if err != nil {
			return
		}
	}

	return
}
