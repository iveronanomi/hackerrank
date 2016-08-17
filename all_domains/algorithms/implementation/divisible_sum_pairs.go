package implementation

import (
	"fmt"
	"math"
)

func DivisibleSumPairs() {
	var n, k int
	fmt.Scan(&n, &k)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}
	var g int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				if d := math.Mod(float64(p[i])+float64(p[j]), float64(k)); d == 0 {
					g++
				}
			}
		}
	}
	fmt.Println(g / 2)
}
