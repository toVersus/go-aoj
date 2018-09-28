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
	N   = 100000
	NIL = -1
)

type Graph struct {
	G     [N][]int
	color [N]int
}

func (g *Graph) dfs(r, c, n int) {
	var S []int
	S = append([]int{r}, S...)
	g.color[r] = c

	for len(S) != 0 {
		u := S[0]
		S = S[1:] // pop
		for i := 0; i < len(g.G[u]); i++ {
			v := g.G[u][i]
			if g.color[v] != NIL {
				continue
			}
			g.color[v] = c
			S = append([]int{v}, S...)
		}
	}
}

func (g *Graph) assignColor(n int) {
	id := 1
	for i := 0; i < n; i++ {
		g.color[i] = NIL
	}

	for u := 0; u < n; u++ {
		if g.color[u] == NIL {
			g.dfs(u, id, n)
			id++
		}
	}
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)

	sc.Scan()
	ss := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(ss[0])
	m, _ := strconv.Atoi(ss[1])

	var g Graph
	for i := 0; i < m; i++ {
		sc.Scan()
		ss := strings.Split(sc.Text(), " ")
		s, _ := strconv.Atoi(ss[0])
		t, _ := strconv.Atoi(ss[1])
		g.G[s] = append(g.G[s], t)
		g.G[t] = append(g.G[t], s)
	}

	g.assignColor(n)

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	var out bytes.Buffer
	for i := 0; i < q; i++ {
		sc.Scan()
		ss := strings.Split(sc.Text(), " ")
		s, _ := strconv.Atoi(ss[0])
		t, _ := strconv.Atoi(ss[1])

		if g.color[s] == g.color[t] {
			out.WriteString("yes\n")
			continue
		}
		out.WriteString("no\n")
	}

	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Print(answer(os.Stdin))
}
