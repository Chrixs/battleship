package shipservice

import (
	"battleship/internal/factory"
	"battleship/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesFleetCorrectly(t *testing.T) {
	ships := CreateNewFleet()

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

	assert.Contains(t, ships, carrier)
	assert.Contains(t, ships, battleship)
	assert.Contains(t, ships, cruiser)
	assert.Contains(t, ships, submarine)
	assert.Contains(t, ships, destroyer)
}

func TestShipsCantOverlap(t *testing.T) {
	carrier := factory.Ship("Carrier", 5)
	carrier.Coordinates = []int{2, 12, 22, 32, 42}
	cruiser := factory.Ship("Cruiser", 3)
	player := factory.Player(1)
	player.Ships = append(player.Ships, carrier, cruiser)

	request := types.DeploymentRequest{
		PlayerId:   1,
		ShipType:   "Cruiser",
		Coordinate: 11,
		IsVertical: false,
	}

	_, err := DeployPlayerShip(request, &player)

	assert.Equal(t, err.Error(), "overlapping ship deployment")
}

func TestShipTypeNotFoundCase(t *testing.T) {
	cruiser := factory.Ship("Cruiser", 3)
	player := factory.Player(1)
	player.Ships = append(player.Ships, cruiser)

	request := types.DeploymentRequest{
		PlayerId:   1,
		ShipType:   "FlyingDoomShip",
		Coordinate: 11,
		IsVertical: false,
	}

	_, err := DeployPlayerShip(request, &player)

	assert.Equal(t, err.Error(), "ship type given does not match types in play")
}
