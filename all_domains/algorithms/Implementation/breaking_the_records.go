package implementation

import (
	"os"
	"fmt"
)

func BreakingTheRecords() {
	var count, best, b, lowest, l, current int
	fmt.Fscanf(os.Stdin, "%d", &count)
	for i := 0; i < count; i++ {
		fmt.Fscanf(os.Stdin, "%d", &current)
		if i == 0 {
			best = current
			lowest = current
			continue
		}
		if current > best {
			best = current
			b++
		}
		if current < lowest {
			lowest = current
			l++
		}
	}
	fmt.Printf("%d %d\n", b, l)
}
