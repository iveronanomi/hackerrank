package search

import (
	"fmt"
	"os"
)

func readInput(input *os.File) [][]int {
	var tasksCount int
	fmt.Fscanf(input, "%v", &tasksCount)
	tasks := make([][]int, tasksCount)
	for i := 0; i < tasksCount; i++ {
		var countInCase int
		fmt.Fscanf(input, "%v", &countInCase)
		tasks[i] = make([]int, countInCase)
		for j := 0; j < countInCase; j++ {
			fmt.Fscanf(input, "%v", &tasks[i][j])
		}
	}
	return tasks
}

func isSumEqual(d1 []int, d2 []int) bool {
	var sum1, sum2 int
	for _, v := range d1 {
		sum1 += v
	}
	for _, v := range d2 {
		sum2 += v
	}
	return sum1 == sum2
}

func SherlockAndArray() {
	input := os.Stdin
	tasks := readInput(input)
	for _, task := range tasks {
		//We need only half of slice
		limit := len(task)/2 + 1
		for i := 0; i <= limit; i++ {
			if isSumEqual(task[:i], task[i+1:]) {
				fmt.Println("YES")
				break
			}
			if i == limit {
				fmt.Println("NO")
			}
		}
	}
}
