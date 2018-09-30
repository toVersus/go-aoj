package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

type Graph [][]int

func (g *Graph) prim(n int) int {
	var (
		u, minv = 0, 0
	)

	d := make([]int, n)
	p := make([]int, n)
	color := make([]int, n)

	for i := 0; i < n; i++ {
		d[i] = INFTY
		p[i] = -1
		color[i] = WHITE
	}
	d[0] = 0

	for {
		minv = INFTY
		u = -1
		for i := 0; i < n; i++ {
			if minv > d[i] && color[i] != BLACK {
				u = i
				minv = d[i]
			}
		}

		if u == -1 {
			break
		}

		color[u] = BLACK
		for v := 0; v < n; v++ {
			if color[v] != BLACK && (*g)[u][v] != INFTY {
				if d[v] > (*g)[u][v] {
					d[v] = (*g)[u][v]
					p[v] = u
					color[v] = GRAY
				}
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		if p[i] != -1 && (*g)[i][p[i]] != INFTY {
			sum += (*g)[i][p[i]]
		}
	}

	return sum
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	g := make(Graph, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sc.Scan()
			e, _ := strconv.Atoi(sc.Text())
			if e == -1 {
				g[i][j] = INFTY
				continue
			}
			g[i][j] = e
		}
	}

	return fmt.Sprint(g.prim(n))
}

func main() {
	fmt.Println(answer(os.Stdin))
}
