package thirty_days_of_code

import (
	"fmt"
	"os"
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

func Day3IntroTtoConditionalStatements() {
	var d uint16
	input := os.Stdin

	fmt.Fscanf(input, "%d", &d)
	if isWeird(d) == true {
		fmt.Print("Weird")
	} else {
		fmt.Print("Not Weird")
	}
}
