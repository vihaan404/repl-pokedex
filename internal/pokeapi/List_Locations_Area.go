package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageURL *string) (PokeApiJson, error) {
	endPoint := "/location"

	fullURL := baseURL + endPoint

	if pageURL != nil {
		fullURL = *pageURL
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokeApiJson{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeApiJson{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return PokeApiJson{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeApiJson{}, err
	}
	parsedBody := PokeApiJson{}

	error := json.Unmarshal(body, &parsedBody)
	if error != nil {
		return PokeApiJson{}, error
	}
	return parsedBody, nil
}
