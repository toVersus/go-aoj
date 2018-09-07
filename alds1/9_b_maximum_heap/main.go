package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

var H int

type Heap []int

func (h Heap) maxHeapify(idx int) {
	var largest int
	l := 2 * idx
	r := 2*idx + 1

	if l <= H && h[l] > h[idx] {
		largest = l
	} else {
		largest = idx
	}
	if r <= H && h[r] > h[largest] {
		largest = r
	}

	if largest != idx {
		h[idx], h[largest] = h[largest], h[idx]
		h.maxHeapify(largest)
	}
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	H, _ = strconv.Atoi(sc.Text())

	A := make(Heap, H+1)
	for i := 1; i <= H; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		A[i] = k
	}

	for i := H / 2; i >= 1; i-- {
		A.maxHeapify(i)
	}

	var out bytes.Buffer
	for i := 1; i <= H; i++ {
		out.WriteString(fmt.Sprintf(" %d", A[i]))
	}

	return out.String()
}

func main() {
	fmt.Println(answer(os.Stdin))
}
