package Implementation

import "fmt"

func step(x1, v1, x2, v2 int) bool {
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
	fmt.Scanf("%v %v", &x1, &v1)
	fmt.Scanf("%v %v", &x2, &v2)
	if step(x1, v1, x2, v2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
