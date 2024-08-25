package playerservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesPlayerWithCorrectId(t *testing.T) {
	player := CreateNewPlayer(1)
	assert.Equal(t, player.ID, 1)
}
