package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var (
	timer                  int
	prenum, parent, lowest []int
	visited                []bool
)

type Graph [][]int

func (g *Graph) dfs(current, prev int) {
	prenum[current], lowest[current] = timer, timer
	timer++

	visited[current] = true

	next := 0
	for i := 0; i < len((*g)[current]); i++ {
		next = (*g)[current][i]
		if !visited[next] {
			parent[next] = current
			g.dfs(next, current)
			lowest[current] = min(lowest[current], lowest[next])
		} else if next != prev {
			lowest[current] = min(lowest[current], prenum[next])
		}
	}
}

func (g *Graph) artPoints(n int) string {
	g.dfs(0, -1)

	ap := make(map[int]struct{}, n)
	np := 0
	for i := 1; i < n; i++ {
		p := parent[i]
		if p == 0 {
			np++
		} else if prenum[p] <= lowest[i] {
			ap[p] = struct{}{}
		}
	}

	if np > 1 {
		ap[0] = struct{}{}
	}

	keys := make([]int, len(ap))
	idx := 0
	for k := range ap {
		keys[idx] = k
		idx++
	}

	sort.Ints(keys)

	var out bytes.Buffer
	for _, k := range keys {
		out.WriteString(fmt.Sprintln(k))
	}

	return out.String()
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	timer = 1
	visited = make([]bool, n)
	prenum = make([]int, n)
	parent = make([]int, n)
	lowest = make([]int, n)

	g := make(Graph, n)

	for i := 0; i < m; i++ {
		sc.Scan()
		s, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		g[s] = append(g[s], t)
		g[t] = append(g[t], s)
	}

	return g.artPoints(n)
}

func main() {
	fmt.Print(answer(os.Stdin))
}
