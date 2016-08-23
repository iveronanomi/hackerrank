package thirty_days_of_code

import "fmt"

func f(n int) int {
	if n == 1 {
		return n
	}
	return f(n-1) + n
}

func Recursion() {
	var n int
	fmt.Scan(&n)
	n = f(n)
	fmt.Print(n)
}
