package arrays

import (
	"fmt"
	"os"
	"strings"
)

func ArrayDS() {
	var length int32
	input := os.Stdin
	fmt.Fscanf(input, "%v", &length)
	arr := make([]string, length)

	//for i := length - 1; i >= 0; i-- {
	//	fmt.Fscanf(input, "%v", &arr[i])
	//}
	for i := int32(0); i < length; i++ {
		fmt.Fscanf(input, "%v", &arr[i])
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Print(strings.Join(arr, " "))
}
