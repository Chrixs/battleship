package playerservice

import (
	"battleship/internal/factory"
	gameservice "battleship/internal/game-service"
	shipservice "battleship/internal/ship-service"
	types "battleship/internal/types"
	"errors"
	"strconv"
)

func CreateNewPlayer(id int, playersTurn bool) types.Player {
	player := factory.Player(id, playersTurn)
	player.Ships = shipservice.CreateNewFleet()
	return player
}

func Fire(firingCoordinate int, attacker *types.Player, defender *types.Player) (types.FiredShot, error) {
	if !attacker.PlayersTurn {
		return factory.FiredShot(""), errors.New("it's not player " + strconv.Itoa(attacker.ID) + "'s turn")
	}

	if attacker.Winner || defender.Winner {
		return factory.FiredShot(""), errors.New("game has ended")
	}

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
		attacker.Winner = true
	}

	attacker.PlayersTurn = false
	defender.PlayersTurn = true

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
