package search

import "fmt"

func readStdin() [][]int {
	var tasksCount int
	fmt.Scanf("%v", &tasksCount)
	tasks := make([][]int, tasksCount)
	for i := 0; i < tasksCount; i++ {
		var countInCase int
		fmt.Scanf("%v", &countInCase)
		tasks[i] = make([]int, countInCase)
		for j := 0; j < countInCase; j++ {
			fmt.Scanf("%v", &tasks[i][j])
		}
	}
	return tasks
}

func isSumEqual(d1 []int, d2 []int) bool {
	var s1, s2 int
	for _, v := range d1 {
		s1 += v
	}
	for _, v := range d2 {
		s2 += v
	}
	return s1 == s2
}

func SherlockAndArray() {
	tasks := readStdin()
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
