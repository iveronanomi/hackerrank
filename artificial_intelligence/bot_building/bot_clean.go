package bot_building

import (
	"fmt"
)

func BotClean() {
	var x, y int
	var v string
	fmt.Scanf("%d %d", &y, &x)
	var garbage [][2]int
	for i := 0; i < 5; i++ {
		fmt.Scanf("%v", &v)
		for idx, v := range v {
			if string(v) == "d" {
				garbage = append(garbage, [2]int{idx, i})
			}
		}
	}
	fmt.Print(next([2]int{x,y}, garbage))
}

func next(current [2]int, list [][2]int) string {
	var closest [2]int
	closestRange := float64(10)

	for _, point := range list {
		if r := heuristic(point, current); r < closestRange {
			closestRange = r
			closest = point
		}
	}

	if current == closest {
		return "CLEAN"
	}
	if current[0] < closest[0] && current[1] == closest[1] {
		return "RIGHT"
	}
	if current[0] > closest[0] && current[1] == closest[1] {
		return "LEFT"
	}
	if current[0] == closest[0] && current[1] < closest[1] {
		return "DOWN"
	}
	if current[0] == closest[0] && current[1] > closest[1] {
		return "UP"
	}
	if current[0] < closest[0] && current[1] < closest[1] {
		return "RIGHT"
	}
	if current[0] > closest[0] && current[1] > closest[1] {
		return "LEFT"
	}
	if current[0] > closest[0] && current[1] < closest[1] {
		return "DOWN"
	}
	if current[0] < closest[0] && current[1] > closest[1] {
		return "UP"
	}
	return "WHOOPS"
}