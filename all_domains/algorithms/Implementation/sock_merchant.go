package implementation

import (
	"fmt"
)

func SockMerchant() {
	var count, color, pairs, n int
	fmt.Scanf("%d", &count)
	pail := make(map[int]int, count)
	for n < count {
		fmt.Scan(&color)
		pail[color]++
		n++
	}
	for _, v := range pail {
		pairs = pairs + int(v/2)
	}
	fmt.Print(pairs)
}
