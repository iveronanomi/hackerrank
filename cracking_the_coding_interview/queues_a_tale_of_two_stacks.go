package cracking_the_coding_interview

import "fmt"

func QueuesATaleOfTwoStacks() {
	var c, t int
	fmt.Scan(&c)
	queue := &Queue{}
	for i := 0; i < c; i++ {
		fmt.Scan(&t)
		switch t {
		case 1: {
			fmt.Scan(&t)
			queue.Enqueue(t)
		}
		case 2: {
			queue.Dequeue()
		}
		case 3: {
			queue.Print()
		}
		}
	}
}

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(n int) {
	q.elements = append(q.elements, n)
}

func (q *Queue) Dequeue() int {
	e := q.elements[0]
	q.elements = q.elements[1:]
	return e
}

func (q *Queue) Print() {
	fmt.Println(q.elements[0])
}