package bot_building

import (
	"fmt"
	"strings"
)

func BotSavesPrincess() {
	var n int
	var v string
	var start, finish [2]int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &v)
		if idx := strings.Index(v, "m"); idx >= 0 {
			start = [2]int{idx,i}
		}
		if idx := strings.Index(v, "p"); idx >= 0 {
			finish = [2]int{idx,i}
		}
	}
	//fmt.Println(start, finish)
	path := find(start, finish, n)
	//fmt.Println(path)
	for _, v := range path {
		fmt.Println(v)
	}
}

func find(start, finish [2]int, n int) []string {
	queue := &Queue{}
	queue.Push(start)
	path := map[[2]int][2]int{start:[2]int{-1,-1}}

	for queue.Len() > 0 {
		current := queue.Pop()
		if current[0] == finish[0] && current[1] == finish[1] {
			break
		}
		for _, next := range neigh(current, n) {
			if _, ok := path[next]; !ok {
				path[next] = current
				queue.Push(next)
			}
		}
	}

	return interpreter(reconstruct(path, start, finish))
}

func reconstruct(cameFrom map[[2]int][2]int, start, finish [2]int) [][2]int{
	current := finish
	path := [][2]int{}

	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)
	//for i, j:= 0, len(path) -1; i < j; i, j = i+1, j -1 {
	//    path[i], path[j] = path[j], path[i]
	//}
	return path
}


func neigh(point [2]int, n int) [][2]int {
	neighbours := [][2]int{
		[2]int{point[0], point[1] - 1},
		[2]int{point[0] + 1, point[1]},
		[2]int{point[0], point[1] + 1},
		[2]int{point[0] - 1, point[1]},
	}
	var result [][2]int
	for _, v := range neighbours {
		if v[0] < 0 || v[0] > (n - 1) || v[1] < 0 || v[1] > (n - 1) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func interpreter(path [][2]int) []string {
	var result []string
	for idx, point := range path {
		if idx == 0 {
			continue
		}
		if point[0] - 1 == path[idx-1][0] && point[1] == path[idx-1][1] {
			result = append(result,"LEFT")
		}
		if point[0] + 1 == path[idx-1][0] && point[1] == path[idx-1][1] {
			result = append(result,"RIGHT")
		}
		if point[0] == path[idx-1][0] && point[1] == path[idx-1][1] + 1 {
			result = append(result,"UP")
		}
		if point[0] == path[idx-1][0] && point[1] == path[idx-1][1] - 1 {
			result = append(result,"DOWN")
		}
	}
	return result
}

type Queue struct {
	elements [][2]int
}

func (q *Queue) Len() int {
	return len(q.elements)
}

func (q *Queue) Pop() [2]int {
	el := q.elements[0]
	q.elements = q.elements[1:]
	return el
}

func (q *Queue) Push(el [2]int) {
	q.elements = append(q.elements, el)
}
