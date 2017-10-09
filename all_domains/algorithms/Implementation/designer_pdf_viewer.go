package implementation

import (
	"fmt"
	"os"
)

func DesignerPDFViewer() {
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	heights := map[string]int{}
	var h int

	for _, l := range chars {
		fmt.Fscanf(os.Stdin, "%d", &h)
		heights[l] = h
	}
	var word string
	fmt.Fscanf(os.Stdin, "%s", &word)
	var maxHeight int
	for i := 0; i < len(word); i++ {
		c := heights[string(word[i])]
		if c > maxHeight {
			maxHeight = c
		}
		if maxHeight == 7 {
			break
		}
	}

	fmt.Println(maxHeight * len(word))
}
