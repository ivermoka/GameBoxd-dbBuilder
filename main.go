package main

import (
	"fmt"
	"ivermoka/GameBoxd-dbBuilder/lib"
)

func main() {
	err := lib.Init()
	if err != nil {
		fmt.Println("Error initializing client: ", err)
	}
}