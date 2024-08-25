package shipservice

import (
	gameservice "battleship/internal/game-service"
	"battleship/internal/types"
	"errors"
	"slices"
)

func CreateNewFleet() []types.Ship {
	carrier := types.Ship{
		Type:   "Carrier",
		Length: 5,
		Health: 5,
	}

	battleship := types.Ship{
		Type:   "Battleship",
		Length: 4,
		Health: 4,
	}

	cruiser := types.Ship{
		Type:   "Cruiser",
		Length: 3,
		Health: 3,
	}

	submarine := types.Ship{
		Type:   "Submarine",
		Length: 3,
		Health: 3,
	}

	destroyer := types.Ship{
		Type:   "Destroyer",
		Length: 2,
		Health: 2,
	}

	return []types.Ship{carrier, battleship, cruiser, submarine, destroyer}
}

func DeployPlayerShip(request types.DeploymentRequest, player *types.Player) (*types.Player, error) {
	var occupiedGridSpaces []int
	var deployingShipIndex *int

	for index, ship := range player.Ships {
		if ship.Type == request.ShipType {
			deployingShipIndex = &index
		} else {
			occupiedGridSpaces = append(occupiedGridSpaces, ship.GridSpaces...)
		}
	}

	if deployingShipIndex == nil {
		return nil, errors.New("ship type given does not match types in play")
	}

	deployedShip, err := gameservice.DeployShip(request.GridNumber, request.IsVertical, player.Ships[*deployingShipIndex])
	if err != nil {
		return nil, err
	}

	for _, gridSpace := range deployedShip.GridSpaces {
		if slices.Contains(occupiedGridSpaces, gridSpace) {
			return nil, errors.New("overlapping ship deployment")
		}
	}

	player.Ships[*deployingShipIndex] = deployedShip
	return player, nil
}
