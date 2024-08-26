package playerservice

import (
	"battleship/internal/factory"
	types "battleship/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesPlayerWithCorrectId(t *testing.T) {
	player := CreateNewPlayer(1, true)
	assert.Equal(t, player.ID, 1)
}

func TestCantFireInSamePlaceTwice(t *testing.T) {
	attacker := CreateNewPlayer(1, true)
	attacker.ShotsFired = []int{100}
	defender := CreateNewPlayer(2, false)

	_, err := Fire(100, &attacker, &defender)
	assert.Equal(t, err.Error(), "already fired at this location")
}

func TestReturnsWinner(t *testing.T) {
	ship := factory.Ship("Smallship", 1)
	ship.Coordinates = []int{1}
	ship.Health = 1

	playerOne := factory.Player(1, true)
	playerTwo := factory.Player(2, false)
	playerTwo.Ships = []types.Ship{ship}

	firedShot, _ := Fire(1, &playerOne, &playerTwo)

	assert.Equal(t, firedShot.Winner, true)
}

func TestPlayerCantFireOutOfTurn(t *testing.T) {
	playerOne := factory.Player(1, false)
	playerTwo := factory.Player(2, true)

	_, err := Fire(1, &playerOne, &playerTwo)

	assert.Equal(t, err.Error(), "it's not player 1's turn")
}

func TestPlayerCantAfterGameWin(t *testing.T) {
	playerOne := factory.Player(1, true)
	playerTwo := factory.Player(2, false)

	playerTwo.Winner = true

	_, err := Fire(1, &playerOne, &playerTwo)

	assert.Equal(t, err.Error(), "game has ended")
}
