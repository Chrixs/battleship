package factory

import (
	"battleship/internal/types"
)

func Ship(shipType string, length int) types.Ship {
	return types.Ship{
		Type:   shipType,
		Length: length,
		Health: length,
	}
}

func Player(id int, playersTurn bool) types.Player {
	return types.Player{
		ID:          id,
		PlayersTurn: playersTurn,
	}
}

func FiredShot(shotStatus string) types.FiredShot {
	return types.FiredShot{
		Status: shotStatus,
	}
}
