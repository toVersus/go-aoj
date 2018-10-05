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
	INFTY = 1 << 32
)

type Graph [][]int

func (g *Graph) floyd(n int) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if (*g)[i][k] == INFTY {
				continue
			}
			for j := 0; j < n; j++ {
				if (*g)[k][j] == INFTY {
					continue
				}
				(*g)[i][j] = min((*g)[i][j], (*g)[i][k]+(*g)[k][j])
			}
		}
	}
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
	e, _ := strconv.Atoi(sc.Text())

	d := make(Graph, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				d[i][j] = 0
				continue
			}
			d[i][j] = INFTY
		}
	}

	for i := 0; i < e; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		c, _ := strconv.Atoi(sc.Text())
		d[u][v] = c
	}

	d.floyd(n)

	negative := false
	for i := 0; i < n; i++ {
		if d[i][i] < 0 {
			negative = true
		}
	}

	if negative {
		return fmt.Sprint("NEGATIVE CYCLE")
	}

	var out bytes.Buffer
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != 0 {
				out.WriteString(" ")
			}
			if d[i][j] == INFTY {
				out.WriteString("INF")
				continue
			}
			out.WriteString(fmt.Sprint(d[i][j]))
		}
		out.WriteString("\n")
	}

	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
