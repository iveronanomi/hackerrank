package implementation

import (
	"fmt"

	"github.com/ereminIvan/hackerrank/utils"
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

func Kangaroo(getInput utils.GetInput) {
	var x1, x2, v1, v2 int
	input := getInput()
	fmt.Fscanf(input, "%v %v", &x1, &v1)
	fmt.Fscanf(input, "%v %v", &x2, &v2)
	if kangarooStep(x1, v1, x2, v2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
