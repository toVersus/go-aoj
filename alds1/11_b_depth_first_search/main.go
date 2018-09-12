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

const N = 100

const (
	WHITE = iota
	GRAY
	BLACK
)

type Graph struct {
	M               [N][N]int
	color, d, f, nt [N]int
	tt              int
}

func (g *Graph) next(u, n int) int {
	for v := g.nt[u]; v < n; v++ {
		g.nt[u] = v + 1
		if g.M[u][v] != 0 {
			return v
		}
	}
	return -1
}

func (g *Graph) dfsVisit(r, n int) {
	var S []int
	S = append([]int{r}, S...)
	g.color[r] = GRAY
	g.tt++
	g.d[r] = g.tt

	for len(S) != 0 {
		u := S[0]
		v := g.next(u, n)
		if v == -1 {
			S = S[1:]
			g.color[u] = BLACK
			g.tt++
			g.f[u] = g.tt
			continue
		}

		if g.color[v] != WHITE {
			continue
		}

		g.color[v] = GRAY
		g.tt++
		g.d[v] = g.tt
		S = append([]int{v}, S...)
	}
}

func (g *Graph) dfs(n int) string {
	for i := 0; i < n; i++ {
		g.color[i] = WHITE
	}
	g.tt = 0

	for u := 0; u < n; u++ {
		if g.color[u] == WHITE {
			g.dfsVisit(u, n)
		}
	}

	var out bytes.Buffer
	for i := 0; i < n; i++ {
		out.WriteString(fmt.Sprintf("%d %d %d\n", i+1, g.d[i], g.f[i]))
	}

	return strings.TrimRight(out.String(), "\n")
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var g Graph
	for i := 0; i < n; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		u--
		for j := 0; j < k; j++ {
			sc.Scan()
			v, _ := strconv.Atoi(sc.Text())
			v--
			g.M[u][v] = 1
		}
	}

	return g.dfs(n)
}

func main() {
	fmt.Println(answer(os.Stdin))
}
