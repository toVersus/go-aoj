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

const INFTY = 1 << 30

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

func (h Heap) extract() int {
	if H < 1 {
		return -INFTY
	}
	maxv := h[1]
	h[1] = h[H]
	H--
	h.maxHeapify(1)
	return maxv
}

func (h *Heap) increaseKey(idx, key int) {
	(*h)[idx] = key
	for idx > 1 && (*h)[idx/2] < (*h)[idx] {
		(*h)[idx], (*h)[idx/2] = (*h)[idx/2], (*h)[idx]
		idx = idx / 2
	}
}

func (h *Heap) insert(key int) {
	H++
	*h = append(*h, -INFTY)
	h.increaseKey(H, key)
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	A := &Heap{0}
	var out bytes.Buffer
	for {
		sc.Scan()
		com := sc.Text()
		switch com {
		case "insert":
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			A.insert(k)
		case "extract":
			out.WriteString(fmt.Sprintf("%d\n", A.extract()))
		case "end":
			return strings.TrimRight(out.String(), "\n")
		}
	}
}

func main() {
	fmt.Println(answer(os.Stdin))
}
