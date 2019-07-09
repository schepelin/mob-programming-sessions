package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GameOfLife(t *testing.T) {
	initState := [5][5]int8{
		[5]int8{0, 0, 0, 0, 0},
		[5]int8{0, 0, 1, 0, 0},
		[5]int8{0, 0, 0, 1, 0},
		[5]int8{0, 1, 1, 1, 0},
		[5]int8{0, 0, 0, 0, 0},
	}
	expectedState := [5][5]int8{
		[5]int8{0, 0, 0, 0, 0},
		[5]int8{0, 0, 0, 0, 0},
		[5]int8{0, 1, 0, 1, 0},
		[5]int8{0, 0, 1, 1, 0},
		[5]int8{0, 0, 1, 0, 0},
	}
	l := NewLife(initState)

	assert.Equal(t, expectedState, l.Tick())
}
