package cracking_the_coding_interview

import (
	"fmt"
	"sort"
)

// todo: Terminated due to timeout
func HeapsFindTheRunningMedian() {
	var c int
	fmt.Scan(&c)
	var m, d float64
	var data []float64
	for i := 0; i < c; i++ {
		fmt.Scan(&d)
		data = append(data, d)
		sort.Sort(sort.Float64Slice(data))
		if i == 0 {
			fmt.Printf("%.01f\n", d)
			continue
		}
		if i % 2 == 0 {
			m = data[i/2]
		} else {
			m = (data[i/2] + data[i/2 + 1])/2
		}
		fmt.Printf("%.01f\n", m)
	}
}
