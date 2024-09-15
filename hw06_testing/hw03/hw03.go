package hw03

import (
	// "fmt" .
	"errors"
	"strings"
)

func HW03(size int) (string, error) {
	// size := 8
	var board strings.Builder
	// fmt.Println("Enter the board size... (8 is the default one)")
	// fmt.Scanf("%d \n", &size)
	if size <= 0 {
		// fmt.Println("Invalid board size.")
		return "", errors.New("invalid board size")
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
				board.WriteString("#")
			} else {
				board.WriteString(" ")
			}
		}
		board.WriteString("\n")
	}
	// fmt.Printf("Printing %d x %d chess board...\n", size, size)
	// fmt.Println(board.String())
	return board.String(), nil
}