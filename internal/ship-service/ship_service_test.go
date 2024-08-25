package shipservice

import (
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
