package baseball

import (
	"fmt"
	log "github.com/apex/log"
	"type-registry/players"
)

type Players struct {
	Players []Player `json:"players"`
}

type Player struct {
	Player players.PlayerBase `json:"player"`
	Hits   int                `json:"hits"`
	Ab     int                `json:"ab"`
}

func (p *Player) PlayerToString() string {
	battingAverage := float32(p.Hits) / float32(p.Ab)
	return fmt.Sprintf("%s is %d years old. Plays for %s. Hits %.3f", p.Player.Name, p.Player.Age, p.Player.Team, battingAverage)
}

func (p *Player) Score() {
	log.Infof("The %s is at the plate.", p.Player.Position)
	log.Infof("%s swings and gets a hit!", p.Player.Name)
	log.Infof("He's rounding first and headed over to second.")
	log.Infof("SAFE AT SECOND and headed to third.")
	log.Infof("Throw gets by the 3rd Baseman and he's headed home!")
	log.Infof("%s SCORES!", p.Player.Name)
}
