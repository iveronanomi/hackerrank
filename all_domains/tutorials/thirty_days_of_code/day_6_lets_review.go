package thirty_days_of_code

import (
	"fmt"
)

func Day6LetsReview() {
	var c int
	var str string
	fmt.Scan(&c)
	for i := 0; i < c; i++ {
		fmt.Scan(&str)
		var odd, even string
		for idx, v := range str {
			if idx%2 == 0 {
				odd = odd + string(v)
			} else {
				even = even + string(v)
			}
		}
		fmt.Printf("%s %s\n", odd, even)
	}
}
