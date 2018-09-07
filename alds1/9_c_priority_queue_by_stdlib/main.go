package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	var x int
	size := len(*h)
	x, *h = (*h)[size-1], (*h)[:size-1]
	return x
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	A := &IntHeap{}
	heap.Init(A)

	var out bytes.Buffer
	for {
		sc.Scan()
		com := sc.Text()
		switch com {
		case "insert":
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			heap.Push(A, k)
		case "extract":
			out.WriteString(fmt.Sprintf("%d\n", heap.Pop(A)))
		case "end":
			return strings.TrimRight(out.String(), "\n")
		}
	}
}

func main() {
	fmt.Println(answer(os.Stdin))
}
