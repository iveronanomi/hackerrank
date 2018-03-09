package cracking_the_coding_interview

import (
	"fmt"
)

func HashTablesRansomNote() {
	var m, l int
	fmt.Scanf("%d %d", &m, &l)

	magazine := scanToMap(m)
	letter := scanToMap(l)

	for k, v := range letter {
		if c, ok := magazine[k]; !ok || v > c {
			fmt.Print("No")
			return
		}
	}
	fmt.Print("Yes")
}

func scanToMap(m int) map[string]int {
	result := map[string]int{}
	var n int
	var w string
	for n < m {
		fmt.Scan(&w)
		if _, ok := result[w]; !ok {
			result[w] = 0
		}
		result[w]++
		n++
	}
	return result
}