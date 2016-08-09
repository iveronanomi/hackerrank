package strings

import "fmt"

//SOSSPSSQSSOR 3 SOSSOT 1
func MarsExploration() {
	var d int
	var input string
	fmt.Scanf("%v", &input)
	for i := 0; len(input)/3 > i; i++ {
		c := input[3*i : 3*i+3]
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
