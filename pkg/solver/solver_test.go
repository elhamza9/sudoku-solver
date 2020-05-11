package solver


import "testing"
func TestSolve(t *testing.T) {
	// 1. Test solvable grid
	solvable := [SIZE][SIZE]int{
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
	solved, _ := solve(solvable);
	if solved == false {
		t.Error("Error Solving solvable puzzle");
	}

	// 2. Test unsolvable grid
	unsolvable := [SIZE][SIZE]int {
		{ 5, 1, 6, 8, 4, 9, 7, 3, 2 },
		{ 3, 0, 7, 6, 0, 5, 0, 0, 0 },
		{ 8, 0, 9, 7, 0, 0, 0, 6, 5 },
		{ 1, 3, 5, 0, 6, 0, 9, 0, 7 },
		{ 4, 7, 2, 5, 9, 1, 0, 0, 6 },
		{ 9, 6, 8, 3, 7, 0, 0, 5, 0 },
		{ 2, 5, 3, 1, 8, 6, 0, 7, 4 },
		{ 6, 8, 4, 2, 0, 7, 5, 0, 0 },
    	{ 7, 9, 1, 0, 5, 0, 6, 0, 8 },
	};

	solved, _ = solve(unsolvable);
	if solved {
		t.Error("Error solving unsolvable puzzle")
	}
}