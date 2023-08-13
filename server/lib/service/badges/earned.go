package badges

import (
	"database/sql"
	"errors"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"math"
	"time"
)

type EarnedBadge struct {
	Id       sqlserver.UUID `json:"id,omitempty"`
	UserId   sqlserver.UUID `json:"user_id,omitempty"`
	BadgeId  sqlserver.UUID `json:"badge_id,omitempty"`
	EarnedAt string         `json:"earned_at,omitempty"`
}

const earnedBadgeCreateTableQuery = `CREATE TABLE [dbo].[EarnedBadges] (
    [id]        UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_EarnedBadges_id] DEFAULT (newid()) NOT NULL,
    [user_id]   UNIQUEIDENTIFIER NOT NULL,
    [badge_id]  UNIQUEIDENTIFIER NOT NULL,
    [earned_at] DATETIME2 (7)    NOT NULL,
    CONSTRAINT [PK_EarnedBadges] PRIMARY KEY CLUSTERED ([id] ASC)
);`

type AddEarnedBadgeRequest struct {
	BadgeID string `json:"badge_id"`
	UserID  string `json:"user_id"`
}

func AddEarnedBadge(container sqlserver.Container, userId sqlserver.UUID, badgeId sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`INSERT INTO [dbo].[EarnedBadges] ([user_id], [badge_id], [earned_at]) VALUES (@P1, @P2, @P3)`)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(sql.Named("userId", userId), sql.Named("badgeId", badgeId), sql.Named("earnedAt", time.Now()))
	return
}

type GetEarnedBadgesRequest struct {
	UserID string `query:"user_id" required:"true"`
}

type GetEarnedBadgeResponse struct {
	Id       string `json:"id"`
	UserID   string `json:"user_id"`
	BadgeID  string `json:"badge_id"`
	EarnedAt string `json:"earned_at"`
}

func GetEarnedBadgesByUserId(container sqlserver.Container, userId sqlserver.UUID) (badges []EarnedBadge, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`SELECT [id], [user_id], [badge_id], [earned_at] FROM [dbo].[EarnedBadges] WHERE [user_id] = @P1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var badge EarnedBadge
		err = rows.Scan(&badge.Id, &badge.UserId, &badge.BadgeId, &badge.EarnedAt)
		if err != nil {
			return
		}
		badges = append(badges, badge)
	}

	return
}

type CountUsersEarnedBadgeRequest struct {
	BadgeID string `query:"badge_id" required:"true"`
}

type CountUsersEarnedBadgeResponse struct {
	Count int `json:"count"`
}

func CountUsersEarnedBadge(container sqlserver.Container, badgeId sqlserver.UUID) (count int, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`SELECT COUNT(*) FROM [dbo].[EarnedBadges] WHERE [badge_id] = @P1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(badgeId)
	if err != nil {
		return
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&count)
	return
}

func CountEarnedBadges(container sqlserver.Container, userId sqlserver.UUID) (count int, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`SELECT COUNT(*) FROM [dbo].[EarnedBadges] WHERE [user_id] = @P1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&count)
	return
}

func GetEarnedRateOfBadge(container sqlserver.Container, badgeId sqlserver.UUID) (rate float64, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`SELECT COUNT(*) FROM [dbo].[EarnedBadges] WHERE [badge_id] = @P1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(badgeId)
	if err := row.Err(); err != nil {
		return rate, err
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		return
	}

	stmt, err = client.Prepare(`SELECT COUNT(*) FROM [dbo].[Users]`)
	if err != nil {
		return
	}
	defer stmt.Close()

	row = stmt.QueryRow()
	if err := row.Err(); err != nil {
		return rate, err
	}

	var total int
	err = row.Scan(&total)
	if err != nil {
		return
	}

	rate = math.Floor((float64(count) / float64(total)) * 100)
	return
}

func CheckEarnedBadge(container sqlserver.Container, userId sqlserver.UUID, badgeId sqlserver.UUID) (result bool, date string, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`SELECT [earned_at] FROM [dbo].[EarnedBadges] WHERE [user_id] = @P1 AND [badge_id] = @P2`)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId, badgeId)
	if err := row.Err(); err != nil {
		return result, date, err
	}

	var earnedAt string
	err = row.Scan(&earnedAt)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}

	result = earnedAt != ""
	date = earnedAt
	return
}
