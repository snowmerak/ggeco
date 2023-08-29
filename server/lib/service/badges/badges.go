package badges

import (
	"github.com/snowmerak/ggeco/server/gen/bean"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
)

type Badge struct {
	Id            sqlserver.UUID `json:"id,omitempty"`
	Name          string         `json:"name,omitempty"`
	Summary       string         `json:"summary,omitempty"`
	ActiveImage   string         `json:"active_image,omitempty"`
	InactiveImage string         `json:"inactive_image,omitempty"`
	SelectedImage string         `json:"selected_image,omitempty"`
}

type GetBadgeRequest struct {
	BadgeID string `query:"id" required:"true"`
}

type GetBadgeResponse struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Summary       string `json:"summary,omitempty"`
	ActiveImage   string `json:"active_image,omitempty"`
	InactiveImage string `json:"inactive_image,omitempty"`
	SelectedImage string `json:"selected_image,omitempty"`
}

func Get(container sqlserver.Container, id sqlserver.UUID) (result Badge, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [name], [summary], [active_image], [inactive_image], [selected_image] from [dbo].[Badges] WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Err(); err != nil {
		return result, err
	}

	if err := row.Scan(&result.Id, &result.Name, &result.Summary, &result.ActiveImage, &result.InactiveImage, &result.SelectedImage); err != nil {
		return result, err
	}

	return
}

type GetBadgeByNameRequest struct {
	BadgeName string `query:"name"`
}

//func GetByName(container sqlserver.Container, name string) (result []Badge, err error) {
//	client, err := sqlserver.GetClient(container)
//	if err != nil {
//		return
//	}
//
//	stmt, err := client.Prepare("SELECT [id], [name], [summary], [image] from [dbo].[Badges] WHERE [name] = @P1")
//	if err != nil {
//		return
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(name)
//	if err != nil {
//		return
//	}
//
//	for rows.Next() {
//		var badge Badge
//		err = rows.Scan(&badge.Id, &badge.Name, &badge.Summary, &badge.Image)
//		if err != nil {
//			return
//		}
//		result = append(result, badge)
//	}
//
//	return
//}

func GetList(container sqlserver.Container) (result []Badge, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [name], [summary], [active_image], [inactive_image], [selected_image] from [dbo].[Badges]")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var badge Badge
		err = rows.Scan(&badge.Id, &badge.Name, &badge.Summary, &badge.ActiveImage, &badge.InactiveImage, &badge.SelectedImage)
		if err != nil {
			return
		}
		result = append(result, badge)
	}

	return
}

func GetSearchables(container bean.Container) (result []Badge, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [name], [summary], [active_image], [inactive_image], [selected_image] from [dbo].[Badges] WHERE [searchable] = 1")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var badge Badge
		err = rows.Scan(&badge.Id, &badge.Name, &badge.Summary, &badge.ActiveImage, &badge.InactiveImage, &badge.SelectedImage)
		if err != nil {
			return
		}
		result = append(result, badge)
	}

	return
}

type AddBadgeRequest struct {
	Name          string `json:"name" required:"true"`
	Summary       string `json:"summary" required:"true"`
	ActiveImage   string `json:"active_image" required:"true"`
	InactiveImage string `json:"inactive_image" required:"true"`
	SelectedImage string `json:"selected_image" required:"true"`
	Searchable    bool   `json:"searchable" required:"true"`
}

type AddBadgeResponse struct {
	Id string `json:"id,omitempty"`
}

func Add(container sqlserver.Container, name string, summary string, activeImage string, inactiveImage string, selectedImage string, searchable bool) (id sqlserver.UUID, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare(`DECLARE @InsertedId TABLE (id UNIQUEIDENTIFIER)
INSERT INTO [dbo].[Badges] ([name], [summary], [active_image], [inactive_image], [selected_image], [searchable]) OUTPUT INSERTED.[id] INTO @InsertedId VALUES (@P1, @P2, @P3, @P4, @P5, @P6)
SELECT id from @InsertedId
`)
	if err != nil {
		return
	}

	row := stmt.QueryRow(name, summary, activeImage, inactiveImage, selectedImage, searchable)
	if err := row.Err(); err != nil {
		return id, err
	}

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return
}

type UpdateBadgeNameRequest struct {
	BadgeID string `json:"badge_id" required:"true"`
	Name    string `json:"name" required:"true"`
}

func UpdateName(container sqlserver.Container, id sqlserver.UUID, name string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Badges] SET [name] = @P2 WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name)

	return
}

type UpdateBadgeSummaryRequest struct {
	BadgeID string `json:"badge_id" required:"true"`
	Summary string `json:"summary" required:"true"`
}

func UpdateSummary(container sqlserver.Container, id sqlserver.UUID, summary string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Badges] SET [summary] = @P2 WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, summary)

	return
}

type UpdateBadgeImageRequest struct {
	BadgeID       string `json:"badge_id" required:"true"`
	ActiveImage   string `json:"image" required:"true"`
	InactiveImage string `json:"image" required:"true"`
	SelectedImage string `json:"image" required:"true"`
}

func UpdateImage(container sqlserver.Container, id sqlserver.UUID, activeImage string, inactiveImage string, selectedImage string) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("UPDATE [dbo].[Badges] SET [active_image] = @P2, [inactive_image] = @P3, [selected_image] = @P4 WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, activeImage, inactiveImage, selectedImage)

	return
}

type DeleteBadgeRequest struct {
	BadgeID string `query:"id" required:"true"`
}

func Delete(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("DELETE FROM [dbo].[Badges] WHERE [id] = @P1")
	if err != nil {
		return
	}

	_, err = stmt.Exec(id)

	return
}
