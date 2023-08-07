package courses

import (
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
)

type CourseBadge struct {
	Id       sqlserver.UUID `json:"id"`
	CourseId sqlserver.UUID `json:"course_id"`
	BadgeId  sqlserver.UUID `json:"badge_id"`
}

type GetCourseBadgesRequest struct {
	CourseId string `query:"course_id" required:"true"`
}

type GetCourseBadgeResponse struct {
	BadgeId   string `json:"badge_id"`
	BadgeName string `json:"badge_name"`
}

func GetCourseBadges(container sqlserver.Container, courseId sqlserver.UUID) (result []CourseBadge, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [badge_id] from [dbo].[CourseBadges] WHERE [course_id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(courseId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b CourseBadge
		err = rows.Scan(&b.BadgeId)
		if err != nil {
			return
		}
		result = append(result, b)
	}

	return
}

type SetCourseBadgesRequest struct {
	CourseId string   `json:"course_id" required:"true"`
	BadgeIds []string `json:"badge_ids" required:"true"`
}

func SetCourseBadges(container sqlserver.Container, courseId sqlserver.UUID, badgeIds []sqlserver.UUID) error {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	tx, err := client.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM [dbo].[CourseBadges] WHERE [course_id] = @P1")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	_, err = stmt.Exec(courseId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	stmt, err = tx.Prepare("INSERT INTO [dbo].[CourseBadges] ([course_id], [badge_id]) VALUES (@P1, @P2)")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	for _, badgeId := range badgeIds {
		_, err = stmt.Exec(courseId, badgeId)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}

			return err
		}
	}

	return tx.Commit()
}
