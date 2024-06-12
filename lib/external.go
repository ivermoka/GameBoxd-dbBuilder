package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func FetchGames() ([]Game, error) {
	var allGames []Game
	offset := 0

	for {
		fmt.Println(offset/100, "Fetching with offset ", offset)
		apiURL := fmt.Sprintf("https://www.giantbomb.com/api/games/?api_key=%s&format=json&limit=100&offset=%d", os.Getenv("GIANT_BOMB_API_KEY"), offset)

		resp, err := http.Get(apiURL)
		if err != nil {
			fmt.Println("Failed to fetch game. Continuing. ", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 420 {
			fmt.Println("API rate limit exceeded. Waiting before retrying...")
			time.Sleep(30 * time.Minute)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Unexpected HTTP response status. Continuing. ", resp.Status)
			continue
		}

		var response struct {
			Results []Game `json:"results"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			fmt.Printf("failed to decode JSON response: %v", err)
			continue
		}

		allGames = append(allGames, response.Results...)

		if len(response.Results) < 100 {
			break
		}

		offset += 100
	}

	return allGames, nil
}

