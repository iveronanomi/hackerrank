package bot_building

import (
	"fmt"
	"math"
	"strings"
)

func BotSavesPrincess2() {
	var n int
	var v string
	var start, finish [2]int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d %d", &start[1], &start[0])
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &v)
		if idx := strings.Index(v, "p"); idx >= 0 {
			finish = [2]int{idx, i}
			break
		}
	}
	fmt.Println(nextMove(start, finish, n))

}

func nextMove(start, finish [2]int, n int) string {
	bestRange := float64(n)
	bestMove := ""

	moves := map[string][2]int{
		"LEFT":  [2]int{start[0] - 1, start[1]},
		"RIGHT": [2]int{start[0] + 1, start[1]},
		"UP":    [2]int{start[0], start[1] - 1},
		"DOWN":  [2]int{start[0], start[1] + 1},
	}

	for move, point := range moves {
		if point[0] < 0 || point[0] > (n-1) || point[1] < 0 || point[1] > (n-1) {
			continue
		}
		if h := heuristic(finish, point); h < bestRange {
			bestRange = h
			bestMove = move
		}
	}

	return bestMove
}

func heuristic(p1, p2 [2]int) float64 {
	return math.Abs(float64(p1[0])-float64(p2[0])) + math.Abs(float64(p1[1])-float64(p2[1]))
}
