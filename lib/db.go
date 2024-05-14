package lib

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

type Image struct {
	IconURL        string `json:"icon_url"`
	MediumURL      string `json:"medium_url"`
	ScreenURL      string `json:"screen_url"`
	ScreenLargeURL string `json:"screen_large_url"`
	SmallURL       string `json:"small_url"`
	SuperURL       string `json:"super_url"`
	ThumbURL       string `json:"thumb_url"`
	TinyURL        string `json:"tiny_url"`
}

type Platform struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Game struct {
	Name             string   `json:"name"`
	DateAdded        string   `json:"date_added"`
	DateLastUpdated  string   `json:"date_last_updated"`
	Deck             string   `json:"deck"`
	Image            Image    `json:"image"`
	Platforms        []Platform `json:"platforms"`
}

func Handler() error {
	supabase, err := Init()
	if err != nil {
		fmt.Println("Error initializing client: ", err)
	}

	games, err := FetchGames()
	if err != nil {
		return fmt.Errorf("error fetching games: %v", err)
	}

	err = InsertGames(supabase, games)
	if err != nil {
		return fmt.Errorf("error inserting games into DB: %v", err)
	}
	fmt.Println("Games successfully added to database.")

	return nil
}

func Init() (*supa.Client, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	supabase := supa.CreateClient(os.Getenv("API_URL"), os.Getenv("API_KEY"))

	return supabase, nil
}

// type Test struct {
// 	Created_at string `json:"created_at"`
// }

// var results []Test

// row := Test {
// 	Created_at: time.Now().Format(time.RFC3339),
// }

// err := supabase.DB.From("test").Insert(row).Execute(&results)
// if err != nil {
// 	return fmt.Errorf("error inserting into db: %w", err)
// }

func InsertGames(supabase *supa.Client, games []Game) error {
	for _, game := range games {
		err := supabase.DB.From("games").Insert(game)
		if err != nil {
			return fmt.Errorf("error inserting game into database: %v", err)
		}
	}

	return nil
}