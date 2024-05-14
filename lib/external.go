package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


func FetchGames() ([]Game, error) {
	apiURL := fmt.Sprintf("https://www.giantbomb.com/api/games/?api_key=%s&format=json", os.Getenv("GIANT_BOMB_API_KEY"))

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch games: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP response status: %s", resp.Status)
	}

	var response struct {
		Results []Game `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return response.Results, nil
}
