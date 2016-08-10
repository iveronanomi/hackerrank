package strings

import (
	"fmt"
	sp "strings"

	"github.com/ereminIvan/hackerrank/utils"
)

func CamelCase(getInput utils.GetInput) {
	var inputString string
	var count int
	input := getInput()

	fmt.Fscanf(input, "%v", &inputString)
	if len(inputString) > 0 {
		count = 1
	}
	for _, cr := range inputString {
		if string(cr) == sp.ToUpper(string(cr)) {
			count++
		}
	}
	fmt.Println(count)
}
