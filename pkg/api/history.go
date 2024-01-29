package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ImageHistory struct {
	ID        string   `json:"Id"`
	Tags      []string `json:"Tags,omitempty"`
	Created   int64    `json:"Created,omitempty"`
	CreatedBy string   `json:"CreatedBy,omitempty"`
	Size      int64    `json:"Size,omitempty"`
	Comment   string   `json:"Comment,omitempty"`
}

func (c *Client) GetHistory(imageId string) ([]ImageHistory, error) {
	resp, err := c.do("GET", fmt.Sprintf("http://localhost/images/%s/history", imageId), nil)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, ErrNoSuchImage
		} else {
			return nil, err
		}
	}

	var imageHistories []ImageHistory
	if err := json.NewDecoder(resp.Body).Decode(&imageHistories); err != nil {
		return nil, err
	}

	return imageHistories, nil
}
