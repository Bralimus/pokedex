package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(area string) (Location, error) {
	url := baseURL + "/location-area/" + area

	if cachedData, found := c.cache.Get(url); found {
		locationResp := Location{}
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
