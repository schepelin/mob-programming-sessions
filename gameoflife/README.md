# General

[Wikipedia: Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)

## Principles

- Test-Driven
- Consumer-First
- Refine existing code

## Rules

Cell can be in one of two possible states, alive or dead
Every cell interacts with its eight neighbors, which are the cells that are horizontally, vertically, or diagonally adjacent.

1. Any live cell with fewer than two live neighbors dies, as if by underpopulation.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by overpopulation.
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

Above rules apply simultaneously to every cell in the grid.

Births and deaths occur simultaneously, and the discrete moment at which this happens is sometimes called a tick.
