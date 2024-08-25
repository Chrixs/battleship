package playerservice

import (
	shipservice "battleship/internal/ship-service"
	types "battleship/internal/types"
)

func CreateNewPlayer(id int) types.Player {
	player := types.Player{
		ID:    id,
		Ships: shipservice.CreateNewFleet(),
	}
	return player
}
