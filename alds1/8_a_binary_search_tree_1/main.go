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
	key                 int
	parent, left, right *Node
}

var out bytes.Buffer

func (n *Node) insert(k int) {
	root := n
	var target *Node
	node := &Node{key: k, left: nil, right: nil}

	for root != nil {
		target = root
		if root.key > k {
			root = root.left
			continue
		}
		root = root.right
	}

	if target.key > node.key {
		target.left = node
		return
	}
	target.right = node
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
			out.WriteString("\n")
		}
	}

	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
