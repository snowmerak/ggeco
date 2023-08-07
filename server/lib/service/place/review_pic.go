package place

import (
	"errors"
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"sort"
)

type ReviewPicture struct {
	Id           sqlserver.UUID `json:"id"`
	ReviewId     sqlserver.UUID `json:"review_id"`
	Order        int            `json:"order"`
	PictureUrl   string         `json:"picture_url"`
	ThumbnailUrl string         `json:"thumbnail_url"`
}

type GetReviewPicturesRequest struct {
	ReviewId string `query:"review_id" required:"true"`
}

type GetReviewPicturesResponse struct {
	Pictures   []string `json:"pictures"`
	Thumbnails []string `json:"thumbnails"`
}

func GetReviewPictures(container sqlserver.Container, reviewId sqlserver.UUID) (result []ReviewPicture, err error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return
	}

	stmt, err := client.Prepare("SELECT [id], [review_id], [order], [picture_url], [thumbnail_url] from [dbo].[PlaceReviewPictures] WHERE [review_id] = @P1 ORDER BY [order]")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(reviewId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p ReviewPicture
		err = rows.Scan(&p.Id, &p.ReviewId, &p.Order, &p.PictureUrl, &p.ThumbnailUrl)
		if err != nil {
			return
		}
		result = append(result, p)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Order < result[j].Order
	})

	return
}

type SetReviewPicturesRequest struct {
	ReviewId      string   `query:"review_id" required:"true"`
	PictureUrls   []string `query:"picture_urls" required:"true"`
	ThumbnailUrls []string `query:"thumbnail_urls" required:"true"`
}

func SetReviewPictures(container sqlserver.Container, reviewId sqlserver.UUID, pictureUrls []string, thumbnailUrls []string) error {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	if len(pictureUrls) != len(thumbnailUrls) {
		return errors.New("picture_urls and thumbnail_urls must have the same length")
	}

	tx, err := client.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM [dbo].[PlaceReviewPictures] WHERE [review_id] = @P1")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(reviewId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	stmt, err = tx.Prepare("INSERT INTO [dbo].[PlaceReviewPictures] ([review_id], [order], [picture_url], [thumbnail_url]) VALUES (@P1, @P2, @P3, @P4)")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	for i, pictureUrl := range pictureUrls {
		_, err = stmt.Exec(reviewId, i, pictureUrl, thumbnailUrls[i])
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}

			return err
		}
	}

	return tx.Commit()
}
