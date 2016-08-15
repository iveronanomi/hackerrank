package thirty_days_of_code

import (
	"fmt"
)

func Day5Loops() {
	var d int

	fmt.Scan(&d)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d", d, i, d*i)
	}
}
