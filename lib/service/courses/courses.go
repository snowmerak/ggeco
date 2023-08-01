package courses

import "github.com/snowmerak/ggeco/lib/client/sqlserver"

type Course struct {
	Id       sqlserver.UUID `json:"id,omitempty"`
	AuthorID sqlserver.UUID `json:"author_id,omitempty"`
	Name     string         `json:"name,omitempty"`
	RegDate  string         `json:"reg_date,omitempty"`
	Review   string         `json:"review,omitempty"`
}

func Delete(container sqlserver.Container, id sqlserver.UUID) (err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("DELETE FROM [dbo].[Courses] WHERE [id] = @P1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return
}
