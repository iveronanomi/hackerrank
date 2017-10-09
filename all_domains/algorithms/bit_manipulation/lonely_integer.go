package bit_manipulation

import "fmt"

func LonelyInteger() {
	var current, unique, count int

	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		fmt.Scan(&current)
		// if not exist in unique list
		if unique&current == 0 {
			// add to unique
			unique |= current
		} else {
			// remove from unique
			unique ^= current
		}
	}

	fmt.Println(unique)
}
