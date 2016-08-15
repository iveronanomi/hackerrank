package thirty_days_of_code

import "fmt"

func Day7Arrays() {
	var c int
	fmt.Scan(&c)
	ds := make([]int, c)
	for i := 0; i < c; i++ {
		fmt.Scan(&ds[i])
	}
	for i := c - 1; i >= 0; i-- {
		fmt.Printf("%d ", ds[i])
	}
}
