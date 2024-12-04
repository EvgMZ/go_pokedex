package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (PokeApiLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	val, ok := c.cache.Get(url)
	if ok {
		locationsPokeApi := PokeApiLocations{}
		err := json.Unmarshal(val, &locationsPokeApi)
		if err != nil {
			return PokeApiLocations{}, err
		}
		return locationsPokeApi, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiLocations{}, err
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return PokeApiLocations{}, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return PokeApiLocations{}, err
	}
	locationsPokeApi := PokeApiLocations{}
	err = json.Unmarshal(data, &locationsPokeApi)
	if err != nil {
		return PokeApiLocations{}, err
	}
	c.cache.Add(url, data)
	return locationsPokeApi, nil

}

func (c *Client) GetLocation(location string) (Location, error) {
	url := baseUrl + "/location-area/" + location
	val, ok := c.cache.Get(url)
	if ok {
		locationPokeApi := Location{}
		err := json.Unmarshal(val, &locationPokeApi)
		if err != nil {
			return Location{}, err
		}
		return locationPokeApi, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Location{}, err
	}
	locationsPokeApi := Location{}
	err = json.Unmarshal(data, &locationsPokeApi)
	if err != nil {
		return Location{}, err
	}
	c.cache.Add(url, data)
	return locationsPokeApi, nil
}
