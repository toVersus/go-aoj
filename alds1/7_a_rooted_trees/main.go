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

type Node struct {
	parent, right, left int
}

type Deps []int

func main() {
	fmt.Println(answer(os.Stdin))
}

func answer(reader io.Reader) string {
	var v, d, c, l, r int

	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	A := make([]Node, n)
	for i := 0; i < n; i++ {
		A[i].parent, A[i].right, A[i].left = -1, -1, -1
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		d, _ = strconv.Atoi(sc.Text())
		for j := 0; j < d; j++ {
			sc.Scan()
			c, _ = strconv.Atoi(sc.Text())
			if j == 0 {
				A[v].left = c
			} else {
				A[l].right = c
			}
			l = c
			A[c].parent = v
		}
	}
	for i := 0; i < n; i++ {
		if A[i].parent == -1 {
			r = i
		}
	}

	D := make(Deps, n)
	rec(A, D, r, 0)

	var out bytes.Buffer
	for i := 0; i < n; i++ {
		out.WriteString(String(A, D, i) + "\n")
	}

	return strings.TrimRight(out.String(), "\n")
}

func rec(T []Node, D Deps, u, p int) int {
	D[u] = p
	if T[u].right != -1 {
		rec(T, D, T[u].right, p)
	}
	if T[u].left != -1 {
		rec(T, D, T[u].left, p+1)
	}

	return 0
}

func String(T []Node, D Deps, u int) string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("node %d: ", u))
	out.WriteString(fmt.Sprintf("parent = %d, ", T[u].parent))
	out.WriteString(fmt.Sprintf("depth = %d, ", D[u]))

	if T[u].parent == -1 {
		out.WriteString("root, ")
	} else if T[u].left == -1 {
		out.WriteString("leaf, ")
	} else {
		out.WriteString("internal node, ")
	}

	out.WriteString("[")

	c := T[u].left
	for {
		if c == -1 {
			break
		}
		out.WriteString(fmt.Sprintf("%d, ", c))
		c = T[c].right
	}

	return strings.TrimRight(out.String(), ", ") + "]"
}
