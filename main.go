package main

import (
	"time"
	"github.com/Warbo360/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands: getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
