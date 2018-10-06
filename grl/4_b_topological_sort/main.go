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
	INFTY = 1 << 29
)

var (
	out   []int
	V     []bool
	indeg []int
)

type Graph [][]int

func (g *Graph) bfs(s int) {
	var q []int
	q = append(q, s)
	V[s] = true
	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		out = append(out, u)
		for i := 0; i < len((*g)[u]); i++ {
			v := (*g)[u][i]
			indeg[v]--
			if indeg[v] == 0 && !V[v] {
				V[v] = true
				q = append(q, v)
			}
		}
	}
}

func (g *Graph) tsort(n int) string {
	for i := 0; i < n; i++ {
		indeg[i] = 0
	}

	for u := 0; u < n; u++ {
		for i := 0; i < len((*g)[u]); i++ {
			v := (*g)[u][i]
			indeg[v]++
		}
	}

	for u := 0; u < n; u++ {
		if indeg[u] == 0 && !V[u] {
			g.bfs(u)
		}
	}

	var buf bytes.Buffer
	for _, o := range out {
		buf.WriteString(fmt.Sprintln(o))
	}

	return strings.TrimRight(buf.String(), "\n")
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	V = make([]bool, n)
	indeg = make([]int, n)
	out = []int{}

	g := make(Graph, n)

	for i := 0; i < m; i++ {
		sc.Scan()
		s, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		g[s] = append(g[s], t)
	}

	return g.tsort(n)
}

func main() {
	fmt.Println(answer(os.Stdin))
}
