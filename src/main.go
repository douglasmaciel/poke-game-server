package main

import (
	"fmt"

	"douglasmaciel.com/poke-game-server/types"
)

func main() {
	p1, err := types.NewPlayer("Catia")
	if err != nil {
		fmt.Println(err)
	}

	id := p1.ID.String()

	p2, err := types.NewPlayerWithId(id, "Catia Nunes")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1 == p2)
	fmt.Println(p1.IsEqual(p2))
}
