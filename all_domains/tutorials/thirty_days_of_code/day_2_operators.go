package thirty_days_of_code

import (
	"fmt"

	"github.com/ereminIvan/hackerrank/utils"
)

func Day2Operators(getInput utils.GetInput) {
	var mealCost float64
	var tipPercent, taxPercent int64
	input := getInput()

	fmt.Fscanf(input, "%f\n%d\n%d", &mealCost, &tipPercent, &taxPercent)
	totalCost := mealCost * float64(100+tipPercent+taxPercent) * .01

	fmt.Printf("The total meal cost is %.f dollars.", totalCost)
}
