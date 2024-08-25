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

func GetPlayerFromId(playerOne *types.Player, playerTwo *types.Player, id int) *types.Player {
	if playerOne.ID == id {
		return playerOne
	}
	return playerTwo
}
