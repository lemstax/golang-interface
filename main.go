package main

import (
	"github.com/apex/log"
	"type-registry/players/factory"
)

func main() {
	ballPlayers, err := factory.UnmarshalBallPlayer("baseball", "./examples")
	if err != nil {
		panic(err)
	}
	fballPlayers, err := factory.UnmarshalBallPlayer("football", "./examples")
	if err != nil {
		panic(err)
	}
	for _, player := range fballPlayers {
		ballPlayers = append(ballPlayers, player)
	}

	for _, player := range ballPlayers {
		log.Info(player.PlayerToString())
		player.Score()
	}
}
