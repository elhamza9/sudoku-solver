package sudoku

// Grid size: SIZE x SIZE
const SIZE int = 9;


// InitGrid provides the initial Grid to be solved.
// Empty cells to be solved for are filled with zero.
func InitGrid() (initialGrid [SIZE][SIZE]int) {
	initialGrid = [SIZE][SIZE]int{
		{ 0, 9, 4, 0, 3, 0, 1, 0, 0 },
		{ 8, 1, 2, 7, 0, 0, 0, 9, 6 },
		{ 3, 0, 0, 1, 9, 0, 0, 0, 0 },
		{ 0, 3, 0, 9, 0, 4, 6, 0, 0 },
		{ 0, 0, 8, 6, 1, 3, 0, 4, 9 },
		{ 0, 0, 6, 2, 0, 0, 0, 0, 1 },
		{ 4, 0, 3, 5, 0, 0, 0, 0, 8 },
		{ 5, 0, 0, 0, 2, 0, 7, 0, 0 },
    	{ 0, 6, 0, 0, 0, 8, 4, 1, 5 },
	};

	return initialGrid;
}

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