package gameservice

import (
	"battleship/internal/types"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestItDeploys(t *testing.T) {
	ship := types.Ship{
		Type:   "Smallship",
		Length: 1,
	}

	deployedShip, _ := DeployShip(1, false, ship)
	assert.Equal(t, deployedShip.GridSpaces, []int{1})
}

func TestItCantDeployAlreadyDeployedShip(t *testing.T) {
	ship := types.Ship{
		Type:       "Smallship",
		Length:     1,
		GridSpaces: []int{1},
	}

	_, err := DeployShip(1, false, ship)
	assert.Equal(t, err.Error(), "cannot deploy ship that is already deployed")
}

func TestItDeploysHorizontally(t *testing.T) {
	ship := types.Ship{
		Type:   "Destroyer",
		Length: 2,
	}

	deployedShip, _ := DeployShip(1, false, ship)
	assert.Equal(t, deployedShip.GridSpaces, []int{1, 2})
}

func TestDeployCantOverflowRow(t *testing.T) {
	ship := types.Ship{
		Type:   "Cruiser",
		Length: 3,
	}

	_, err := DeployShip(9, false, ship)
	assert.Equal(t, err.Error(), "ship deployment exceeds game bounds")
}

func TestItDeploysVertically(t *testing.T) {
	ship := types.Ship{
		Type:   "Destroyer",
		Length: 2,
	}

	deployedShip, _ := DeployShip(1, true, ship)
	assert.Equal(t, deployedShip.GridSpaces, []int{1, 11})
}

func TestDeployCantExceedColumnBounds(t *testing.T) {
	ship := types.Ship{
		Type:   "Cruiser",
		Length: 3,
	}

	_, err := DeployShip(91, true, ship)
	assert.Equal(t, err.Error(), "ship deployment exceeds game bounds")
}

// func TestShipsCantOverlap(t *testing.T) {
// 	carrier := types.Ship{
// 		Type:   "Carrier",
// 		Length: 5,
// 	}

// 	DeployShip(2, true, carrier)

// 	cruiser := types.Ship{
// 		Type:   "Cruiser",
// 		Length: 3,
// 	}
// 	_, err := DeployShip(11, false, cruiser)

// 	assert.Equal(t, err.Error(), "overlapping ship deployment")
// }
