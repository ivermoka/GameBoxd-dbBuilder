package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

type Test struct {
	Created_at string `json:"created_at"`
}


func Init() error {
	if err := godotenv.Load(); err != nil {
        return fmt.Errorf("error loading .env file: %w", err)
    }
	supabase := supa.CreateClient(os.Getenv("API_URL"), os.Getenv("API_KEY"))

	var results []Test

	row := Test {
		Created_at: time.Now().Format(time.RFC3339),
	}

	err := supabase.DB.From("test").Insert(row).Execute(&results)
	if err != nil {
		return fmt.Errorf("error inserting into db: %w", err)
	}

	err = supabase.DB.From("test").Select("*").Execute(&results)
	if err != nil {
		return fmt.Errorf("error selecting from db: %w", err)
	}

	fmt.Println(results)
	return nil
}
