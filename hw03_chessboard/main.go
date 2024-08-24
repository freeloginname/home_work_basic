package main

import (
	"fmt"
)

func main() {
	size := 8
	var board string
	var cellCharacter string
	fmt.Println("Enter the board size... (8 is the default one)")
	fmt.Scanf("%d \n", &size)
	if size < 0 {
		fmt.Println("Only natural numbers supported.")
		return
	}
	for line := 0; line < size; line++ {
		var oddCondition int
		if line%2 == 1 {
			oddCondition = 1
		} else {
			oddCondition = 0
		}
		for cell := 0; cell < size; cell++ {
			if cell%2 == oddCondition {
				cellCharacter = "#"
			} else {
				cellCharacter = " "
			}

			board += cellCharacter
		}
		board += "\n"
	}
	fmt.Printf("Printing %d x %d chess board...\n", size, size)
	fmt.Println(board)
}
