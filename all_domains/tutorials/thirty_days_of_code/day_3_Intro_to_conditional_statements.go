package thirty_days_of_code

import (
	"fmt"

	"github.com/ereminIvan/hackerrank/utils"
)

func isWeird(d uint16) bool {
	if d%2 != 0 {
		return true
	}
	if d >= 6 && d <= 20 {
		return true
	}
	if d >= 2 && d <= 5 || d > 20 {
		return false
	}
	return true
}

func Day3IntroTtoConditionalStatements(getInput utils.GetInput) {
	input := getInput()
	var d uint16
	fmt.Fscanf(input, "%d", &d)
	if isWeird(d) == true {
		fmt.Print("Weird")
	} else {
		fmt.Print("Not Weird")
	}
}
