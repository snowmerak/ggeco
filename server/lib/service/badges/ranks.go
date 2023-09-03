package badges

import (
	"database/sql"
	"errors"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
)

/*
CREATE TABLE [dbo].[BadgeRank] (
    [user_id]      UNIQUEIDENTIFIER NOT NULL,
    [current_rank] BIGINT           NOT NULL,
    [prev_rank]    BIGINT           NULL,
    [update_at]    DATETIME2 (7)    NULL,
    CONSTRAINT [PK_BadgeRank] PRIMARY KEY CLUSTERED ([user_id] ASC)
);
*/

func GetMyRank(container bean.Container, userId sqlserver.UUID) (rank int64, delta int64, updated string, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [current_rank], [prev_rank], [update_at] FROM [dbo].[BadgeRank] WHERE [user_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	updatedAt := sql.NullString{}
	prevRank := sql.NullInt64{}
	err = stmt.QueryRow(userId).Scan(&rank, &prevRank, &updatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
		rank = 0
	}

	if prevRank.Valid {
		delta = rank - prevRank.Int64
	}

	if updatedAt.Valid {
		updated = updatedAt.String
	}

	return
}
