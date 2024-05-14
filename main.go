package main

import (
	"fmt"
	"ivermoka/GameBoxd-dbBuilder/lib"
)

func main() {
	err := lib.Handler()
	if err != nil {
		fmt.Println("Error running handler: ", err)
	}
}