package strings

import (
	"fmt"
	sp "strings"
)

func CamelCase () {
	var input string
	var count int
	fmt.Scanf("%v", &input)
	if len(input) > 0 {
		count = 1;
	}
	for _, cr := range input {
		if string(cr) == sp.ToUpper(string(cr)) {
			count++
		}
	}
	fmt.Println(count)
}