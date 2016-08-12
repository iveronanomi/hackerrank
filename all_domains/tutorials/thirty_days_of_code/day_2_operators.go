package thirty_days_of_code

import (
	"fmt"
	"os"
)

func Day2Operators() {
	var mealCost float64
	var tipPercent, taxPercent int64
	input := os.Stdin

	fmt.Fscanf(input, "%f\n%d\n%d", &mealCost, &tipPercent, &taxPercent)
	totalCost := mealCost * float64(100+tipPercent+taxPercent) * .01

	fmt.Printf("The total meal cost is %.f dollars.", totalCost)
}
