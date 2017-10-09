package implementation

import (
	"fmt"
	"os"
)

//GradingStudents solving of GradingStudents task
func GradingStudents() {
	var count, grade int
	fmt.Fscanf(os.Stdin, "%d", &count)
	for i := 0; i < count; i++ {
		fmt.Fscanf(os.Stdin, "%d", &grade)
		if grade < 38 {
			fmt.Println(grade)
			continue
		}
		if rmd := grade % 5; rmd >= 3 {
			fmt.Println(grade + 5 - rmd)
			continue
		}
		fmt.Println(grade)
	}
}
