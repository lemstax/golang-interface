package players

type Player interface {
	PlayerToString() string
	Score()
}

type PlayerBase struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Team     string `json:"team"`
	Position string `json:"position"`
}
