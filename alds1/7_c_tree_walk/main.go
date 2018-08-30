package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

const NIL = -1

type Node struct {
	p, l, r int
}

type Nodes []Node

var out bytes.Buffer

func (n Nodes) preParse(u int) {
	if u == NIL {
		return
	}
	out.WriteString(fmt.Sprintf(" %d", u))
	n.preParse(n[u].l)
	n.preParse(n[u].r)
}

func (n Nodes) inParse(u int) {
	if u == NIL {
		return
	}
	n.inParse(n[u].l)
	out.WriteString(fmt.Sprintf(" %d", u))
	n.inParse(n[u].r)
}

func (n Nodes) postParse(u int) {
	if u == NIL {
		return
	}
	n.postParse(n[u].l)
	n.postParse(n[u].r)
	out.WriteString(fmt.Sprintf(" %d", u))
}

func answer(reader io.Reader) string {
	var v, l, r, root int
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nodes := make(Nodes, n)
	for i := 0; i < n; i++ {
		nodes[i].p = NIL
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		l, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())

		nodes[v].l, nodes[v].r = l, r
		if l != NIL {
			nodes[l].p = v
		}
		if r != NIL {
			nodes[r].p = v
		}
	}

	for i := 0; i < n; i++ {
		if nodes[i].p == NIL {
			root = i
		}
	}

	out.WriteString(fmt.Sprintln("Preorder"))
	nodes.preParse(root)
	out.WriteString("\n")
	out.WriteString(fmt.Sprintln("Inorder"))
	nodes.inParse(root)
	out.WriteString("\n")
	out.WriteString(fmt.Sprintln("Postorder"))
	nodes.postParse(root)

	return out.String()
}

func main() {
	fmt.Println(answer(os.Stdin))
}
