package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Node struct {
	key                 int
	parent, left, right *Node
}

var (
	out bytes.Buffer
)

func (n *Node) insert(k int) {
	x := n
	var y *Node
	z := &Node{
		key:   k,
		left:  &Node{},
		right: &Node{},
	}

	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
			continue
		}
		x = x.right
	}

	z.parent = y
	if y == nil {
		n = z
		return
	}
	if z.key < y.key {
		y.left = z
		return
	}
	y.right = z
}

func (n *Node) inorder() {
	if n == nil {
		return
	}
	n.left.inorder()
	if n.key != 0 {
		out.WriteString(fmt.Sprintf(" %d", n.key))
	}
	n.right.inorder()
}

func (n *Node) preorder() {
	if n == nil {
		return
	}
	if n.key != 0 {
		out.WriteString(fmt.Sprintf(" %d", n.key))
	}
	n.left.preorder()
	n.right.preorder()
}

func answer(reader io.Reader) string {
	root := &Node{}

	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		com := sc.Text()
		if com == "insert" {
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			root.insert(k)
		} else if com == "print" {
			root.inorder()
			out.WriteString("\n")
			root.preorder()
		}
	}

	return out.String()
}

func main() {
	fmt.Println(answer(os.Stdin))
}
