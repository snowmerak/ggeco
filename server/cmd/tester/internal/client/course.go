package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/snowmerak/ggeco/server/function/app"
	"strconv"
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

func (c *Client) GetRecentlyCourses(at string, req app.GetRecentCoursesRequest) (app.GetRecentCoursesResponse, error) {
	resp, err := c.Get("/course/recent", at, c.Query("count", strconv.Itoa(req.Count)))
	if err != nil {
		return app.GetRecentCoursesResponse{}, err
	}

	var result app.GetRecentCoursesResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return app.GetRecentCoursesResponse{}, err
	}

	return result, nil
}

func (c *Client) GetCourseInfo(at string, req app.GetCourseInfoRequest) (app.GetCourseInfoResponse, error) {
	data, err := c.Get("/course", at, c.Query("course_id", req.CourseId))
	if err != nil {
		return app.GetCourseInfoResponse{}, err
	}

	resp := app.GetCourseInfoResponse{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return app.GetCourseInfoResponse{}, err
	}

	return resp, nil
}
