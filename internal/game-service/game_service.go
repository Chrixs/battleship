package gameservice

import (
	"battleship/internal/types"
	"errors"
	"slices"
)

func DeployShip(gridNumber int, isVertical bool, ship types.Ship) (types.Ship, error) {
	if ship.GridSpaces != nil {
		return ship, errors.New("cannot deploy ship that is already deployed")
	}

	var deploymentGridSpaces []int
	stop := false

	for i := 0; i < ship.Length; i++ {
		if stop || gridNumber > 100 {
			return ship, errors.New("ship deployment exceeds game bounds")
		}

		deploymentGridSpaces = append(deploymentGridSpaces, gridNumber)

		if isVertical {
			gridNumber += 10
		} else {

			if slices.Contains(endOfRow(), gridNumber) {
				stop = true
			}
			gridNumber++
		}
	}

	ship.GridSpaces = deploymentGridSpaces
	return ship, nil
}

func endOfRow() []int {
	return []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
}
