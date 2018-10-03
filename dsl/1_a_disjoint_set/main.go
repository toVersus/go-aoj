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

type DisjointSet struct {
	rank []int
	p    []int
}

func New(size int) *DisjointSet {
	ds := &DisjointSet{
		rank: make([]int, size),
		p:    make([]int, size),
	}

	for i := 0; i < size; i++ {
		ds.p[i] = i
		ds.rank[i] = 0
	}

	return ds
}

func (ds *DisjointSet) same(x, y int) bool {
	return ds.findSet(x) == ds.findSet(y)
}

func (ds *DisjointSet) unite(x, y int) {
	ds.link(ds.findSet(x), ds.findSet(y))
}

func (ds *DisjointSet) link(x, y int) {
	if ds.rank[x] > ds.rank[y] {
		ds.p[y] = x
	} else {
		ds.p[x] = y
		if ds.rank[x] == ds.rank[y] {
			ds.rank[y]++
		}
	}
}

func (ds *DisjointSet) findSet(x int) int {
	if x != ds.p[x] {
		ds.p[x] = ds.findSet(ds.p[x])
	}
	return ds.p[x]
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	ds := New(n)

	var out bytes.Buffer
	for i := 0; i < q; i++ {
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())

		if t == 0 {
			ds.unite(a, b)
			continue
		} else if ds.same(a, b) {
			out.WriteString("1\n")
			continue
		}
		out.WriteString("0\n")
	}

	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
