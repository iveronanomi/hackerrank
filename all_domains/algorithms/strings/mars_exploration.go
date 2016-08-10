package strings

import (
	"fmt"

	"github.com/ereminIvan/hackerrank/utils"
)

func MarsExploration(getInput utils.GetInput) {
	var d int
	var inputString string
	input := getInput()
	fmt.Fscanf(input, "%v", &inputString)
	for i := 0; len(inputString)/3 > i; i++ {
		c := inputString[3*i : 3*i+3]
		if c[0] != 83 {
			d++
		}
		if c[1] != 79 {
			d++
		}
		if c[2] != 83 {
			d++
		}
	}
	fmt.Println(d)
}
