package solver 

// Grid size: SIZE x SIZE
const SIZE int = 9;


// isSolved checks if the Grid is solved
func isSolved(grid [SIZE][SIZE]int) (solved bool, row int, col int) {
	solved = true;
	for r := 0; r < SIZE; r++ {
		for c := 0; c < SIZE; c++ {
			if grid[r][c] < 1 {
				solved = false;
				row = r;
				col = c;
				return solved, row, col
			}
		}
	}

	return solved, row, col;
}

// isSafe checks if it's safe to fill a cell with a value.
func isSafe(grid [SIZE][SIZE]int, row int, col int, val int) (safe bool) {
	safe = true;

	// 1. Check Row
	for i := 0; i < SIZE; i++ {
		if grid[row][i] == val {
			return false;
		}
	}

	// 2. Check Column
	for i := 0; i < SIZE; i++ {
		if grid[i][col] == val {
			return false;
		}
	}

	// 3. Check Subgrid
	subGridRow := int(row/3)*3;
	subGridCol := int(col/3)*3;
	for r := subGridRow; r < subGridRow+3; r++ {
		for c := subGridCol; c < subGridCol+3; c++ {
			if grid[r][c] == val {
				return false;
			}
		}
	}

	return safe;
}