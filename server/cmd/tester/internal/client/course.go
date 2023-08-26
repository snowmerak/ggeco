package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/snowmerak/ggeco/server/function/app"
)

func (c *Client) AddCourse(at string, req app.SetCourseRequest) (app.SetCourseResponse, error) {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(req)
	if err != nil {
		return app.SetCourseResponse{}, err
	}

	resp, err := c.Post("/course", at, buf)
	if err != nil {
		return app.SetCourseResponse{}, err
	}
	fmt.Println(string(resp))

	var result app.SetCourseResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return app.SetCourseResponse{}, err
	}

	return result, nil
}

func (c *Client) GetFavoriteCourses(at string) (app.GetFavoriteCoursesResponse, error) {
	resp, err := c.Get("/course/favorite", at)
	if err != nil {
		return app.GetFavoriteCoursesResponse{}, err
	}

	var result app.GetFavoriteCoursesResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return app.GetFavoriteCoursesResponse{}, err
	}

	return result, nil
}
