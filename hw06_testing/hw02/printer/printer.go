package printer

import (
	"fmt"
	"strings"

	"github.com/freeloginname/home_work_basic/hw06_testing/hw02/types"
)

func PrintStaff(staff []types.Employee) string {
	var str string
	var allStrings strings.Builder
	for i := 0; i < len(staff); i++ {
		str = fmt.Sprintf(
			"User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		// fmt.Println(str)
		allStrings.WriteString(str)
	}
	return allStrings.String()
}
