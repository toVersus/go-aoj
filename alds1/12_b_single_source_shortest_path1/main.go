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

type Graph [][]int

func (g *Graph) dikstra(n int) string {
	var (
		u, minv = 0, 0
	)

	d := make([]int, n)
	color := make([]int, n)

	for i := 0; i < n; i++ {
		d[i] = INFTY
		color[i] = WHITE
	}
	d[0] = 0
	color[0] = GRAY

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
				if d[v] > d[u]+(*g)[u][v] {
					d[v] = d[u] + (*g)[u][v]
					color[v] = GRAY
				}
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

	g := make(Graph, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = INFTY
		}
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())

		for j := 0; j < k; j++ {
			sc.Scan()
			v, _ := strconv.Atoi(sc.Text())
			sc.Scan()
			c, _ := strconv.Atoi(sc.Text())

			g[u][v] = c
		}
	}

	return g.dikstra(n)
}

func main() {
	fmt.Println(answer(os.Stdin))
}
