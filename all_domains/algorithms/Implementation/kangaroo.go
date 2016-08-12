package implementation

import (
	"fmt"
	"os"
)

func kangarooStep(x1, v1, x2, v2 int) bool {
	for i := 0; i <= 10000; i++ {
		if x1 == x2 {
			return true
		}
		x1 += v1
		x2 += v2
	}
	return false
}

func Kangaroo() {
	var x1, x2, v1, v2 int
	input := os.Stdin
	fmt.Fscanf(input, "%v %v", &x1, &v1)
	fmt.Fscanf(input, "%v %v", &x2, &v2)
	if kangarooStep(x1, v1, x2, v2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
