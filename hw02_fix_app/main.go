package main

import (
	"fmt"

	"main/printer"
	"main/reader"
	"main/types"
)

func main() {
	path := "data.json"
	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)
	var err error
	var staff []types.Employee
	if len(path) == 0 {
		path = "data.json"
	}
	staff, err = reader.ReadJSON(path)
	fmt.Print(err)
	printer.PrintStaff(staff)
}
