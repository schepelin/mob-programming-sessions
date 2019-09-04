package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func emptyGrid(size uint) [][]bool {
	grid := make([][]bool, size)

	for i := range grid {
		grid[i] = make([]bool, size)
	}
	return grid
}

// generates a glider shape in the top left corner
// .#.
// ..#
// ###
func glider(size uint) [][]bool{
	grid := emptyGrid(size)
	grid[0][1] = true
	grid[1][2] = true
	grid[2][0] = true
	grid[2][1] = true
	grid[2][2] = true
	return grid
}

func Test_GameOfLife_Tick(t *testing.T) {
	var gridSize uint = 10
	initialState := glider(gridSize)

	l := NewLife(initialState)

	l.Next()
	expected := `
..........
#.#.......
.##.......
.#........
..........
..........
..........
..........
..........
..........
`

	assert.Equal(t, expected, l.String())
}
