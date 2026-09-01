package main

import (
	"time"
	"github.com/Warbo360/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		commands: getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
