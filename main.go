package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	//fmt.Println(d1, d2)
	var s1, s2 int
	for _, v := range d1 {
		s1 += v
	}
	for _, v := range d2 {
		s2 += v
	}
	return s1 == s2
}

func main() {
	tasks := [][]int{{1, 2, 3}, {1, 2, 3, 3}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 1}}
	b, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	fmt.Printf(string(b))
	for _, task := range tasks {
		var isIt bool
		limit := len(task)
		for i := 1; i < limit; i++ {
			if isSumEqual(task[:i], task[i+1:]) {
				isIt = true
				fmt.Println("YES")
				break
			}
		}
		if !isIt {
			fmt.Println("NO")
		}
	}
}
