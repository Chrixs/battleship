package gameservice

import (
	"battleship/internal/factory"
	"battleship/internal/types"
	"errors"
	"slices"
)

func DeployShip(coordinate int, isVertical bool, ship types.Ship) (types.Ship, error) {
	if ship.Coordinates != nil {
		return ship, errors.New("cannot deploy ship that is already deployed")
	}

	var deploymentCoordinates []int
	stop := false

	for i := 0; i < ship.Length; i++ {
		if stop || coordinate > 100 {
			return ship, errors.New("ship deployment exceeds game bounds")
		}

		deploymentCoordinates = append(deploymentCoordinates, coordinate)

		if isVertical {
			coordinate += 10
		} else {

			if slices.Contains(endOfRow(), coordinate) {
				stop = true
			}
			coordinate++
		}
	}

	ship.Coordinates = deploymentCoordinates
	return ship, nil
}

func FireCalculation(firingCoordinate int, attacker types.Player, defender types.Player) (types.Player, types.FiredShot, error) {
	isHit := false
	for index, ship := range defender.Ships {
		for _, occupiedCoord := range ship.Coordinates {
			if occupiedCoord == firingCoordinate {
				isHit = true
			}
		}
		if isHit {
			defender.Ships[index].Health -= 1
			if defender.Ships[index].Health == 0 {
				firedShot := factory.FiredShot("sunk")
				firedShot.ShipType = ship.Type
				return defender, firedShot, nil
			}
			return defender, factory.FiredShot("hit"), nil
		}
	}

	return defender, factory.FiredShot("miss"), nil
}

func endOfRow() []int {
	return []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
}
