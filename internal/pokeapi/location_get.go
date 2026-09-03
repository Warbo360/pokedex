package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(area string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + area
	if value, ok := c.cache.Get(url); ok {
		locationAreaResp := RespLocationArea{}
		err := json.Unmarshal(value, &locationAreaResp)
		if err != nil {
			return RespLocationArea, err
		}
		return locationAreaResp, nil
	}
}
