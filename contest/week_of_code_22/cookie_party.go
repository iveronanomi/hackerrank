package week_of_code_22

import (
	"fmt"
	"math"
	"os"
)

func CookieParty() {
	var p, c, d float64
	fmt.Fscanf(os.Stdin, "%f %f", &p, &c)
	if p == 0 {
		d = 0
	} else if p > c {
		d = p - c
	} else if p < c {
		d = p*math.Ceil(c/p) - c
	} else {
		d = 0
	}
	fmt.Printf("%.f\n", d)
}
