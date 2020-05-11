package main


import (
	"fmt"
	"sudoku/pkg/solver"
	"strings"
	"strconv"
	"bufio"
	"os"
)


func fillGrid(grid *[solver.SIZE][solver.SIZE]int) {
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter your grid:");
	NewLineLoop:
	for row, line, fields := 0, "", []string{}; row < solver.SIZE; {
		fmt.Printf("\tLine %d: ", row+1);
		line, _ = reader.ReadString('\n');
		fields = strings.Fields(line);
		if len(fields) > 9 {
			fmt.Printf("Line contains more than %d elements. Try again.\n", solver.SIZE);
		} else {
			for col, v := range fields {
				res, err := strconv.Atoi(v);
				if err != nil {
					fmt.Printf("Conversion of %s to int failed. Try again.\n", v);
					continue NewLineLoop; // this will skip incrementing the row.
				} else {
					if res < 0 || res > 9 {
						fmt.Printf("Wrong value %d. Only values from 0 ~ 9 are allowed. Try again.\n", res);
						continue NewLineLoop;
					}
					grid[row][col] = res;
				}
			}
			row++;
		}
	}
}

func printGrid(grid [solver.SIZE][solver.SIZE]int) {
	fmt.Printf("\n\tPuzzle\n");
	for _, row := range grid {
		fmt.Printf("    ")
		for _, cell := range row {
			fmt.Printf("%d ", cell);
		}
		fmt.Printf("\n");
	}
}

func main() {
	var grid [solver.SIZE][solver.SIZE]int;
	fillGrid(&grid);
	printGrid(grid);

	solved, result := solver.Solve(grid);
	if solved {
		fmt.Println("Solved :)");
		printGrid(result);
	} else {
		fmt.Println("Could not solve the given grid :(")
	}
}