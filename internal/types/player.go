package types

type Player struct {
	ID          int    `json:"id"`
	Ships       []Ship `json:"ships"`
	ShotsFired  []int  `json:"shotsFired,omitempty"`
	PlayersTurn bool   `json:"playersTurn"`
	Winner      bool   `json:"winner,omitempty"`
}

type Players struct {
	Players []Player `json:"players"`
}
