package playerservice

import (
	"battleship/internal/factory"
	gameservice "battleship/internal/game-service"
	shipservice "battleship/internal/ship-service"
	types "battleship/internal/types"
	"errors"
)

func CreateNewPlayer(id int) types.Player {
	player := factory.Player(id)
	player.Ships = shipservice.CreateNewFleet()
	return player
}

func Fire(firingCoordinate int, attacker *types.Player, defender *types.Player) (types.FiredShot, error) {
	for _, shot := range attacker.ShotsFired {
		if shot == firingCoordinate {
			return factory.FiredShot(""), errors.New("already fired at this location")
		}
	}

	modifiedDefender, firedShot, err := gameservice.FireCalculation(firingCoordinate, *attacker, *defender)
	if err != nil {
		return factory.FiredShot(""), err
	}

	attacker.ShotsFired = append(attacker.ShotsFired, firingCoordinate)
	*defender = modifiedDefender

	if isDead(defender) {
		firedShot.Winner = true
	}

	return firedShot, nil
}

func GetPlayersFromId(playerOne *types.Player, playerTwo *types.Player, id int) (*types.Player, *types.Player) {
	if playerOne.ID == id {
		return playerOne, playerTwo
	}
	return playerTwo, playerOne
}

func isDead(player *types.Player) bool {
	for _, ship := range player.Ships {
		if ship.Health > 0 {
			return false
		}
	}

	return true
}
