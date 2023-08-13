package place

import (
	"github.com/snowmerak/ggeco/server/lib/client/sqlserver"
	"time"
)

type FavoritePlace struct {
	Id           sqlserver.UUID `json:"id"`
	UserId       sqlserver.UUID `json:"user_id"`
	PlaceId      string         `json:"place_id"`
	RegisteredAt time.Time      `json:"registered_at"`
}

type GetFavoritePlacesByUserIdRequest struct {
	UserId string `query:"user_id"`
}

type GetFavoritePlaceByUserIdResponse struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	PlaceId      string    `json:"place_id"`
	RegisteredAt time.Time `json:"registered_at"`
}

func GetFavoritePlacesByUserId(container sqlserver.Container, userId sqlserver.UUID) ([]FavoritePlace, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return nil, err
	}

	var favoritePlaces []FavoritePlace

	stmt, err := client.Prepare("SELECT [id], [user_id], [place_id], [registered_at] FROM [dbo].[FavoritePlaces] WHERE [user_id] = @P1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var favoritePlace FavoritePlace
		err := rows.Scan(&favoritePlace.Id, &favoritePlace.UserId, &favoritePlace.PlaceId, &favoritePlace.RegisteredAt)
		if err != nil {
			return nil, err
		}
		favoritePlaces = append(favoritePlaces, favoritePlace)
	}

	return favoritePlaces, nil
}

type CountFavoritePlaceRequest struct {
	PlaceId string `query:"place_id"`
}

func CountFavoritePlace(container sqlserver.Container, placeId string) (int, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return 0, err
	}

	stmt, err := client.Prepare("SELECT COUNT(*) FROM [dbo].[FavoritePlaces] WHERE [place_id] = @P1")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(placeId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

type AddFavoritePlaceRequest struct {
	UserId  string `json:"user_id"`
	PlaceId string `json:"place_id"`
}

type AddFavoritePlaceResponse struct {
	Id string `json:"id"`
}

func AddFavoritePlace(container sqlserver.Container, userId sqlserver.UUID, placeId string) (sqlserver.UUID, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return nil, err
	}

	stmt, err := client.Prepare(`DECLARE @insertedId UNIQUEIDENTIFIER
INSERT INTO [dbo].[FavoritePlaces] ([user_id], [place_id], [registered_at]) OUTPUT inserted.id VALUES (@P1, @P2, @P3)
SELECT @insertedId`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId, placeId, time.Now())
	if err := row.Err(); err != nil {
		return nil, err
	}

	var id sqlserver.UUID
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

type DeleteFavoritePlaceRequest struct {
	Id string `query:"id"`
}

func DeleteFavoritePlace(container sqlserver.Container, userId sqlserver.UUID, placeId string) error {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return err
	}

	stmt, err := client.Prepare("DELETE FROM [dbo].[FavoritePlaces] WHERE [user_id] = @P1 AND [place_id] = @P2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, placeId)
	if err != nil {
		return err
	}

	return nil
}

func CheckFavoritePlace(container sqlserver.Container, userId sqlserver.UUID, placeId string) (bool, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return false, err
	}

	stmt, err := client.Prepare("SELECT COUNT(*) FROM [dbo].[FavoritePlaces] WHERE [user_id] = @P1 AND [place_id] = @P2")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(userId, placeId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CountFavoritePlaceByUserId(container sqlserver.Container, userId sqlserver.UUID) (int, error) {
	client, err := sqlserver.GetClient(container)
	if err != nil {
		return 0, err
	}

	stmt, err := client.Prepare("SELECT COUNT(*) FROM [dbo].[FavoritePlaces] WHERE [user_id] = @P1")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
