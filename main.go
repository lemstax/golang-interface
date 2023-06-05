package main

import (
	"github.com/apex/log"
	"type-registry/players"
	"type-registry/players/factory"
)

var ballTypes = []string{"baseball", "football"}

func main() {
	var ballPlayers []players.Player
	for _, ballType := range ballTypes {
		bp, err := factory.UnmarshalBallPlayer(ballType, "./examples")
		if err != nil {
			panic(err)
		}
		ballPlayers = append(ballPlayers, bp...)
	}

	for _, player := range ballPlayers {
		log.Info(player.PlayerToString())
		player.Score()
	}
}
