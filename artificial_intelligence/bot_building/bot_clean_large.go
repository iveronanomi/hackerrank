package bot_building

import (
	"fmt"
)

func BotCleanLarge() {
	var x, y, h, w int
	var v string
	fmt.Scanf("%d %d", &y, &x)
	fmt.Scanf("%d %d", &h, &w)
	var garbage [][2]int
	for i := 0; i < h; i++ {
		fmt.Scanf("%v", &v)
		for idx, v := range v {
			if string(v) == "d" {
				garbage = append(garbage, [2]int{idx, i})
			}
		}
	}
	fmt.Print(next([2]int{x, y}, garbage))
}
