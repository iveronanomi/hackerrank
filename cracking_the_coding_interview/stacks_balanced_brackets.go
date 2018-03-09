package cracking_the_coding_interview

import (
	"fmt"
)

//todo optimize
func StacksBalancedBrackets() {
	var c int
	var w string
	fmt.Scan(&c)
	for i := 0; i < c; i ++ {
		fmt.Scan(&w)
		ok := true
		for ok {
			w, ok = remove(w)
		}
		if len(w) > 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

func remove(s string) (string, bool) {
	for idx := range s {
		if idx == 0 {
			continue
		}
		current := string(s[idx-1]) + string(s[idx])
		if  current == "{}" || current == "[]" || current == "()" {
			return string(s[:idx-1]) + string(s[idx+1:]), true
		}
	}
	return s, false
}