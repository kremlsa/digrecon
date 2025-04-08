package simple

import (
	"fmt"
)

type Grid [3][3]string

var weights [3][3]int = [3][3]int{
	{2, 1, 2},
	{4, -4, 4},
	{2, -1, 2},
}

const bias int = -5

func Recon() {
	fmt.Println("Hello!")
	var grid Grid
	initGrid(&grid)
	enterValues(&grid)
	res := recognition(&grid)
	printGrid(&grid)
	fmt.Printf("I think this number is: %d\n", res)
	fmt.Println("---------------------------------")
}

func recognition(grid *Grid) (result int) {
	sum := 0
	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] == "+" {
				sum += 1 * weights[i][j]
			} else {
				sum += 0 * weights[i][j]
			}
		}
	}
	if sum+bias > 0 {
		result = 0
	} else {
		result = 1
	}
	return result
}

func enterValues(grid *Grid) {
	for i := range len(grid) {
		for j := range len(grid[i]) {
			grid[i][j] = prompt(grid)
		}
	}
}

func prompt(grid *Grid) string {
	var cell string
	printGrid(grid)
Loop:
	for {
		fmt.Println("Please enter value '-' or '+': ")
		fmt.Scanln(&cell)
		if cell != "-" && cell != "+" {
			fmt.Println("Wrong value!")
		} else {
			break Loop
		}
	}
	return cell
}

func printGrid(grid *Grid) {
	for i := range grid {
		fmt.Println(grid[i])
	}
}

func initGrid(grid *Grid) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}
}
