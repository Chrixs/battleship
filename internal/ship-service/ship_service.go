package shipservice

import (
	"battleship/internal/factory"
	gameservice "battleship/internal/game-service"
	"battleship/internal/types"
	"errors"
	"slices"
)

func CreateNewFleet() []types.Ship {
	carrier := factory.Ship("Carrier", 5)
	battleship := factory.Ship("Battleship", 4)
	cruiser := factory.Ship("Cruiser", 3)
	submarine := factory.Ship("Submarine", 3)
	destroyer := factory.Ship("Destroyer", 2)

	return []types.Ship{carrier, battleship, cruiser, submarine, destroyer}
}

func DeployPlayerShip(request types.DeploymentRequest, player *types.Player) (types.Ship, error) {
	var occupiedCoordinates []int
	var deployingShipIndex *int
	var deployedShip types.Ship

	for index, ship := range player.Ships {
		if ship.Type == request.ShipType {
			deployingShipIndex = &index
		} else {
			occupiedCoordinates = append(occupiedCoordinates, ship.Coordinates...)
		}
	}

	if deployingShipIndex == nil {
		return deployedShip, errors.New("ship type given does not match types in play")
	}

	deployedShip, err := gameservice.DeployShip(request.Coordinate, request.IsVertical, player.Ships[*deployingShipIndex])
	if err != nil {
		return deployedShip, err
	}

	for _, coordinate := range deployedShip.Coordinates {
		if slices.Contains(occupiedCoordinates, coordinate) {
			return deployedShip, errors.New("overlapping ship deployment")
		}
	}

	player.Ships[*deployingShipIndex] = deployedShip
	return deployedShip, nil
}
