package football

import (
	"fmt"
	"github.com/apex/log"
	"type-registry/players"
)

type Player struct {
	Player     players.PlayerBase `json:"player"`
	Touchdowns int                `json:"touchdowns"`
}

func (p *Player) PlayerToString() string {
	return fmt.Sprintf("%s is %d years old. Plays %s for %s and has scored %d touchdowns", p.Player.Name, p.Player.Age, p.Player.Position, p.Player.Team, p.Touchdowns)
}

func (p *Player) Score() {
	log.Infof("The %s gets the ball.", p.Player.Position)
	log.Info("He crosses the 40... the 50... 45, 40.")
	log.Info("35... 30.")
	log.Info("25... 20.")
	log.Infof("10... 5... TOUCHDOWN %s", p.Player.Team)
}
