package main

import (
	"fmt"
	"testing"

	"github.com/freeloginname/home_work_basic/hw06_testing/hw02/printer"
	"github.com/freeloginname/home_work_basic/hw06_testing/hw02/reader"
	"github.com/freeloginname/home_work_basic/hw06_testing/hw02/types"
	"github.com/stretchr/testify/require"
)

func TestHW02(t *testing.T) {
	testCases := []struct {
		desc               string
		path               string
		err                error
		readerExpectation  []types.Employee
		printerExpectation string
	}{
		{
			desc: "base case",
			path: "data.json",
			err:  nil,
			readerExpectation: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
			printerExpectation: "User ID: 10; Age: 25; Name: Rob; Department ID: 3; User ID: 11; Age: 30; Name: George; Department ID: 2; ",
		},
		{
			desc: "json missing field case",
			path: "corrupt_data.json",
			err:  nil,
			readerExpectation: []types.Employee{
				{
					UserID:       10,
					Age:          0,
					Name:         "Rob",
					DepartmentID: 3,
				},
			},
			printerExpectation: "User ID: 10; Age: 0; Name: Rob; Department ID: 3; ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			staff, err := reader.ReadJSON(tC.path)
			require.Equal(t, tC.err, err)
			require.Equal(t, tC.readerExpectation, staff)
			require.Equal(t, tC.printerExpectation, printer.PrintStaff(staff))
		})
	}
}

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
