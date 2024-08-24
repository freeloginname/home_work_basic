package main

import (
	"fmt"
)

func main() {
	size := 8
	var board string
	var cell_character string
	fmt.Println("Enter the board size... (8 is the default one)")
	fmt.Scanf("%d \n", &size)
	if size < 0 {
		fmt.Println("Only natural numbers supported.")
		return
	}
	for line := 0; line < size; line = line + 1 {
		var odd_condition int
		if line%2 == 1 {
			odd_condition = 1
		} else {
			odd_condition = 0
		}
		for cell := 0; cell < size; cell = cell + 1 {
			if cell%2 == odd_condition {
				cell_character = "#"
			} else {
				cell_character = " "
			}

			board = board + cell_character
		}
		board = board + "\n"
	}
	fmt.Printf("Printing %d x %d chess board...\n", size, size)
	fmt.Println(board)
}
