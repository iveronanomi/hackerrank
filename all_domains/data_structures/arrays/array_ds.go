package arrays

import (
	"fmt"

	"github.com/ereminIvan/hackerrank/utils"
	"strings"
)

func ArrayDS(getInput utils.GetInput) {
	input := getInput()
	var length int32
	fmt.Fscanf(input, "%v", &length)
	arr := make([]string, length)

	for i := length - 1; i >= 0; i-- {
		fmt.Fscanf(input, "%v", &arr[i])
	}
	fmt.Print(strings.Join(arr, " "))
}
