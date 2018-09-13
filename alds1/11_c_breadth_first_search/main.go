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
	N     = 100
	INFTY = 1 << 21
)

type Graph struct {
	M [N][N]int
	d [N]int
}

func (g *Graph) bfs(s, n int) string {
	var q []int
	q = append(q, s) // push

	for i := 0; i < n; i++ {
		g.d[i] = INFTY
	}
	g.d[s] = 0

	for len(q) != 0 {
		u := q[0]
		q = q[1:] // pop
		for v := 0; v < n; v++ {
			if g.M[u][v] == 0 || g.d[v] != INFTY {
				continue
			}
			g.d[v] = g.d[u] + 1
			q = append(q, v) // push
		}
	}

	var out bytes.Buffer
	for i := 0; i < n; i++ {
		if g.d[i] == INFTY {
			out.WriteString(fmt.Sprintf("%d -1\n", i+1))
			continue
		}
		out.WriteString(fmt.Sprintf("%d %d\n", i+1, g.d[i]))
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

	return g.bfs(0, n)
}

func main() {
	fmt.Println(answer(os.Stdin))
}
