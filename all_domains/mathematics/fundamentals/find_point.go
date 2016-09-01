package fundamentals

import "fmt"

func FindPoint() {
	var c int
	fmt.Scan(&c)
	for ; c > 0; c-- {
		var qx, qy, px, py int
		fmt.Scanf("%d %d %d %d", &qx, &qy, &px, &py)
		fmt.Printf("%d %d\n", 2*px-qx, 2*py-qy)
	}
}
