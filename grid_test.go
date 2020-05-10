package sudoku

import "testing"

func TestInitGrid(t *testing.T) {
	g := InitGrid();

	for r := 0; r < SIZE; r++ {
		for c := 0; c < SIZE; c++ {
			if g[r][c] < 0 || g[r][c] > 9 {
				t.Error("Value Error in initial grid cell")
			}
		}
	}
}

func TestIsSolved(t *testing.T) {
	g := [SIZE][SIZE]int{
		{ 6, 9, 4, 0, 3, 0, 1, 0, 0 },
		{ 8, 1, 2, 7, 0, 0, 0, 9, 6 },
		{ 3, 0, 0, 1, 9, 0, 0, 0, 0 },
		{ 0, 3, 0, 9, 0, 4, 6, 0, 0 },
		{ 0, 0, 8, 6, 1, 3, 0, 4, 9 },
		{ 0, 0, 6, 2, 0, 0, 0, 0, 1 },
		{ 4, 0, 3, 5, 0, 0, 0, 0, 8 },
		{ 5, 0, 0, 0, 2, 0, 7, 0, 0 },
    	{ 0, 6, 0, 0, 0, 8, 4, 1, 5 },
	};
	solved, row, col := isSolved(g);
	if solved == true || row != 0 || col != 3 {
		t.Error("Error isSolved case unsolved grid");
	}

	g = [SIZE][SIZE]int{
		{ 6, 9, 4, 8, 3, 2, 1, 5, 7 },
		{ 8, 1, 2, 7, 4, 5, 3, 9, 6 },
		{ 3, 5, 7, 1, 9, 6, 2, 8, 4 },
		{ 1, 3, 5, 9, 8, 4, 6, 7, 2 },
		{ 7, 2, 8, 6, 1, 3, 5, 4, 9 },
		{ 9, 4, 6, 2, 5, 7, 8, 3, 1 },
		{ 4, 7, 3, 5, 6, 1, 9, 2, 8 },
		{ 5, 8, 1, 4, 2, 9, 7, 6, 3 },
    	{ 2, 6, 9, 3, 7, 8, 4, 1, 5 },
	};
	solved, row, col = isSolved(g);
	if solved == false || row != 0 || col != 0 {
		t.Error("Error isSolved case solved grid");
	}
}

func TestIsSafe(t *testing.T) {
	g := [SIZE][SIZE]int{
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
	if isSafe(g, 0, 0, 1) {
		t.Error("Error isSafe case not safe (value exists in row)");
	}
	if isSafe(g, 0, 0, 8) {
		t.Error("Error isSafe case not safe (value exists in column)");
	}
	if isSafe(g, 0, 0, 2) {
		t.Error("Error isSafe case not safe (value exists in subgrid)");
	}
	if !isSafe(g, 0, 0, 6) {
		t.Error("Error isSafe case safe");
	}
}