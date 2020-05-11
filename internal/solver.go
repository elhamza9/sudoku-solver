package sudoku

// solve recursively solves the given grid
// and returns:
//	- boolean indicating whether it was solved or not
//	- resulting grid
func solve( grid [SIZE][SIZE]int ) (bool, [SIZE][SIZE]int) {

	// 1. Check if Grid is solved
	solved, nextRow, nextCol := isSolved(grid);
	if solved {
		return true, grid;
	}

	// 2. Find the next cell to be filled and solve for it.
	for val := 1; val <= 9; val++ {
		if isSafe(grid, nextRow, nextCol, val) {
			grid[nextRow][nextCol] = val;
			res, resGrid := solve(grid);
			if res {
				return res, resGrid;
			}
		}
	}
	
	return false, grid;
}