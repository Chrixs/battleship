package gameservice

import (
	"battleship/internal/factory"
	"battleship/internal/types"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestItDeploys(t *testing.T) {
	ship := factory.Ship("Smallship", 1)

	deployedShip, _ := DeployShip(1, false, ship)
	assert.Equal(t, deployedShip.Coordinates, []int{1})
}

func TestItCantDeployAlreadyDeployedShip(t *testing.T) {
	ship := factory.Ship("Smallship", 1)
	ship.Coordinates = []int{1}

	_, err := DeployShip(1, false, ship)
	assert.Equal(t, err.Error(), "cannot deploy ship that is already deployed")
}

func TestItDeploysHorizontally(t *testing.T) {
	ship := factory.Ship("Destroyer", 2)

	deployedShip, _ := DeployShip(1, false, ship)
	assert.Equal(t, deployedShip.Coordinates, []int{1, 2})
}

func TestDeployCantOverflowRow(t *testing.T) {
	ship := factory.Ship("Cruiser", 3)

	_, err := DeployShip(9, false, ship)
	assert.Equal(t, err.Error(), "ship deployment exceeds game bounds")
}

func TestItDeploysVertically(t *testing.T) {
	ship := factory.Ship("Destroyer", 2)

	deployedShip, _ := DeployShip(1, true, ship)
	assert.Equal(t, deployedShip.Coordinates, []int{1, 11})
}

func TestDeployCantExceedColumnBounds(t *testing.T) {
	ship := factory.Ship("Cruiser", 3)

	_, err := DeployShip(91, true, ship)
	assert.Equal(t, err.Error(), "ship deployment exceeds game bounds")
}

func TestCanFireOnShipAndReturnHit(t *testing.T) {
	carrier := factory.Ship("Carrier", 5)
	carrier.Coordinates = []int{1, 2, 3, 4, 5}

	playerOne := factory.Player(1, true)
	playerTwo := factory.Player(2, false)
	playerTwo.Ships = []types.Ship{carrier}

	_, firedShot, _ := FireCalculation(1, playerOne, playerTwo)

	assert.Equal(t, firedShot.Status, "hit")
}

func TestCanFireOnShipAndReturnMiss(t *testing.T) {
	carrier := factory.Ship("Carrier", 5)
	carrier.Coordinates = []int{1, 2, 3, 4, 5}

	playerOne := factory.Player(1, true)
	playerTwo := factory.Player(2, false)
	playerTwo.Ships = []types.Ship{carrier}

	_, firedShot, _ := FireCalculation(6, playerOne, playerTwo)

	assert.Equal(t, firedShot.Status, "miss")
}

func TestCanFireOnShipAndDetectSunk(t *testing.T) {
	carrier := factory.Ship("Carrier", 5)
	carrier.Coordinates = []int{1, 2, 3, 4, 5}
	carrier.Health = 1

	playerOne := factory.Player(1, true)
	playerTwo := factory.Player(2, false)
	playerTwo.Ships = []types.Ship{carrier}

	_, firedShot, _ := FireCalculation(1, playerOne, playerTwo)

	assert.Equal(t, firedShot.Status, "sunk")
	assert.Equal(t, firedShot.ShipType, "Carrier")
}
