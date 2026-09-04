package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(area string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + area

	// Check cache first for url and retrieve data if present

	if value, ok := c.cache.Get(url); ok {
		locationAreaResp := RespLocationArea{}
		err := json.Unmarshal(value, &locationAreaResp)
		if err != nil {
			return locationAreaResp, err
		}
		return locationAreaResp, nil
	}

	// If cache misses, build http GET request message and error check

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	// Send request over the network and error check

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}

	// defer body closing from readall io of the response body

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	// Make an empty RespLocationArea struct and unmarshal the json into that struct

	locationAreaResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	// Store the raw data into the cache under the built url

	c.cache.Add(url, dat)

	// return built struct

	return locationAreaResp, nil

}
