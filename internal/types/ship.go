package types

type Ship struct {
	Type       string `json:"type"`
	Length     int    `json:"length"`
	Health     int    `json:"health"`
	GridSpaces []int  `json:"gridSpaces,omitempty"`
}
