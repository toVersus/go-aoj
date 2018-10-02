package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	INFTY = 1 << 21
	NIL   = -1
)

const (
	WHITE = iota
	GRAY
	BLACK
)

type Node struct {
	pair     [2]int
	priority int
	index    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type PriorityQueues []PriorityQueue

func (pq *PriorityQueues) dikstra(n int) string {
	PQ := make(PriorityQueue, n)
	d := make([]int, n)
	color := make([]int, n)

	for i := 0; i < n; i++ {
		d[i] = INFTY
		color[i] = WHITE
	}
	d[0] = 0
	PQ.Push([2]int{0, 0})
	color[0] = GRAY

	for PQ.Len() != 0 {
		f := PQ.Pop()
		u := f.(*Node).pair[1]
		color[u] = BLACK

		if d[u] < f.(*Node).pair[0]*(-1) {
			continue
		}

		for j := 0; j < n; j++ {
			v := (*pq)[u][j].pair[0]
			if color[v] == BLACK {
				continue
			}
			if d[v] > d[u]+(*pq)[u][j].pair[1] {
				d[v] = d[u] + (*pq)[u][j].pair[1]
				PQ.Push([2]int{d[v] * -1, v})
				color[v] = GRAY
			}
		}
	}

	var out bytes.Buffer
	for i := 0; i < n; i++ {
		out.WriteString(fmt.Sprintf("%d ", i))
		if d[i] == INFTY {
			out.WriteString(fmt.Sprintf("%d\n", -1))
			continue
		}
		out.WriteString(fmt.Sprintf("%d\n", d[i]))
	}

	return strings.TrimRight(out.String(), "\n")
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	pq := make(PriorityQueues, n)
	for i := 0; i < n; i++ {
		pq[i] = make(PriorityQueue, n)

		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())

		for j := 0; j < k; j++ {
			sc.Scan()
			v, _ := strconv.Atoi(sc.Text())
			sc.Scan()
			c, _ := strconv.Atoi(sc.Text())

			pq[u].Push([]int{v, c})
		}
	}

	return pq.dikstra(n)
}

func main() {
	fmt.Println(answer(os.Stdin))
}
