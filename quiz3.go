package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

//add game
func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Martin"}

	gamer.AddGame("Petak Umpet")
	gamer.AddGame("Game Kelereng")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
