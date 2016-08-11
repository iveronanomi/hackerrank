package week_of_code_22

import (
	"fmt"
	"math"

	"github.com/ereminIvan/hackerrank/utils"
)

func CookieParty(getInput utils.GetInput) {
	var t [2]float64
	var d float64
	input := getInput()
	fmt.Fscanf(input, "%f %f", &t[0], &t[1])
	if t[0] == 0 {
		d = 0
	} else if t[0] > t[1] {
		d = t[0] - t[1]
	} else if t[0] < t[1] {
		d = t[0]*math.Ceil(t[1]/t[0]) - t[1]
	} else {
		d = 0
	}
	fmt.Printf("%.f\n", d)
}
