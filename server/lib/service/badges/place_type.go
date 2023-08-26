package badges

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"golang.org/x/crypto/blake2s"
)

var _ = `
CREATE TABLE [dbo].[PlaceTypeToBadgeId] (
    [id]         UNIQUEIDENTIFIER NOT NULL,
    [place_type] VARCHAR (64)    NOT NULL,
    [badge_id]   UNIQUEIDENTIFIER NOT NULL,
    CONSTRAINT [PK_PlaceTypeToBadgeId] PRIMARY KEY CLUSTERED ([id] ASC)
);
`

func GetBadgeFromPlaceType(container bean.Container, pt string) (sqlserver.UUID, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return sqlserver.UUID{}, err
	}

	hashed := blake2s.Sum256([]byte(pt))
	pt = hex.EncodeToString(hashed[:])

	stmt, err := client.Prepare("SELECT [badge_id] from [dbo].[PlaceTypeToBadgeId] WHERE [place_type] = @P1")
	if err != nil {
		return sqlserver.UUID{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(pt)
	if err := row.Err(); err != nil {
		return sqlserver.UUID{}, err
	}

	var badgeId sqlserver.UUID
	if err := row.Scan(&badgeId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlserver.UUID{}, nil
		}
		return sqlserver.UUID{}, err
	}

	return badgeId, nil
}

func AddPlaceTypeToBadgeId(container bean.Container, pt string, badgeId sqlserver.UUID) error {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	hashed := blake2s.Sum256([]byte(pt))
	pt = hex.EncodeToString(hashed[:])

	stmt, err := client.Prepare(`INSERT INTO [dbo].[PlaceTypeToBadgeId] ([place_type], [badge_id]) VALUES (@P1, @P2)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sql.Named("placeType", pt), sql.Named("badgeId", badgeId))
	return err
}
