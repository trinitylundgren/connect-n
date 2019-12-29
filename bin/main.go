package main

import (
	"log"

	"github.com/trinitylundgren/connect-n/game"
)

func main() {
	g, err := game.New(6, 7, 4)
	if err != nil {
		log.Fatalf("Could not initiaize game: %v\n", err)
	}
	g.Play()
}
