package imgmnger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

// jetti:bean Client
type ImageManager struct {
	url    string
	client *http.Client
}

func New(url string) *ImageManager {
	return &ImageManager{
		url:    url,
		client: &http.Client{},
	}
}

type UploadResponse struct {
	OriginImageUrl    string `json:"origin_image_url"`
	ThumbnailImageUrl string `json:"thumbnail_image_url"`
}

func (im *ImageManager) Upload(storageName string, userId string, file io.Reader, name string, contentType string, size int) (url UploadResponse, err error) {
	buffer := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buffer)
	part, err := writer.CreateFormFile("image", name)
	if err != nil {
		return url, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return url, err
	}

	uri := fmt.Sprintf("%s&user_id=%s&size=%s&storage_name=%s", im.url, userId, strconv.FormatInt(int64(size), 10), storageName)
	req, err := http.NewRequest(http.MethodPut, uri, buffer)
	if err != nil {
		return url, err
	}

	req.Header.Set("Content-Type", contentType)
	resp, err := im.client.Do(req)
	if err != nil {
		return url, err
	}
	defer resp.Body.Close()

	respMessage, err := io.ReadAll(resp.Body)
	if err != nil {
		return url, err
	}

	if resp.StatusCode != http.StatusOK {
		return url, fmt.Errorf("invalid status code: %d, message: %s", resp.StatusCode, string(respMessage))
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&url)
	if err != nil {
		return url, err
	}

	return url, nil
}
