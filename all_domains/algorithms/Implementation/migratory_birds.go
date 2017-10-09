package implementation

import (
	"fmt"
	"os"
)

func MigratoryBirds() {
	var count, id, max, maxId int
	birds := map[int]int{}

	fmt.Fscanf(os.Stdin, "%d", &count)
	for i := 0; i < count; i++ {
		fmt.Fscanf(os.Stdin, "%d", &id)
		birds[id]++
		if birds[id] > max {
			max = birds[id]
			maxId = id
		}
		if birds[id] == max {
			if id < maxId {
				maxId = id
			}
		}
	}
	fmt.Println(maxId)
}
