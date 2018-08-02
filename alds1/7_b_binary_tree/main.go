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

const NIL = -1

type Node struct {
	parent, left, right int
}

type Deps []int
type Heights []int

type Tree struct {
	T []Node
	D Deps
	H Heights
}

func answer(reader io.Reader) string {
	var v, l, r, root int

	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var tree Tree
	tree.T = make([]Node, n)
	for i := 0; i < n; i++ {
		tree.T[i].parent = NIL
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		l, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())
		tree.T[v].left, tree.T[v].right = l, r
		if l != NIL {
			tree.T[l].parent = v
		}
		if r != NIL {
			tree.T[r].parent = v
		}
	}

	for i := 0; i < n; i++ {
		if tree.T[i].parent == NIL {
			root = i
		}
	}
	tree.D = make(Deps, n)
	tree.H = make(Heights, n)

	tree.setDepth(root, 0)
	tree.setHeight(root)

	var out bytes.Buffer
	for i := 0; i < n; i++ {
		out.WriteString(tree.String(i))
	}

	return strings.TrimRight(out.String(), "\n")
}

func (t *Tree) setDepth(u, d int) {
	if u == NIL {
		return
	}
	t.D[u] = d
	t.setDepth(t.T[u].left, d+1)
	t.setDepth(t.T[u].right, d+1)
}

func (t *Tree) setHeight(u int) int {
	var h1, h2 int
	if t.T[u].left != NIL {
		h1 = t.setHeight(t.T[u].left) + 1
	}
	if t.T[u].right != NIL {
		h2 = t.setHeight(t.T[u].right) + 1
	}

	t.H[u] = h2
	if h1 > h2 {
		t.H[u] = h1
	}
	return t.H[u]
}

func (t *Tree) getSibling(u int) int {
	if t.T[u].parent == NIL {
		return NIL
	}
	if t.T[t.T[u].parent].left != u && t.T[t.T[u].parent].left != NIL {
		return t.T[t.T[u].parent].left
	}
	if t.T[t.T[u].parent].right != u && t.T[t.T[u].parent].right != NIL {
		return t.T[t.T[u].parent].right
	}
	return NIL
}

func (t *Tree) String(u int) string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("node %d: ", u))
	out.WriteString(fmt.Sprintf("parent = %d, ", t.T[u].parent))
	out.WriteString(fmt.Sprintf("sibling = %d, ", t.getSibling(u)))

	var deg int
	if t.T[u].left != NIL {
		deg++
	}
	if t.T[u].right != NIL {
		deg++
	}
	out.WriteString(fmt.Sprintf("degree = %d, ", deg))
	out.WriteString(fmt.Sprintf("depth = %d, ", t.D[u]))
	out.WriteString(fmt.Sprintf("height = %d, ", t.H[u]))

	if t.T[u].parent == NIL {
		out.WriteString("root\n")
	} else if t.T[u].left == NIL && t.T[u].right == NIL {
		out.WriteString("leaf\n")
	} else {
		out.WriteString("internal node\n")
	}

	return out.String()
}

func main() {
	fmt.Println(answer(os.Stdin))
}
