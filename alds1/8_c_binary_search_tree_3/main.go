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

func (n *Node) treeMinimum() *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *Node) find(k int) *Node {
	for n != nil && n.key != k {
		if n.key > k {
			n = n.left
			continue
		}
		n = n.right
	}
	return n
}

func (n *Node) treeSuccessor() *Node {
	if n.right != nil {
		return n.right.treeMinimum()
	}
	t := n.parent
	for t != nil && n == t.right {
		n = t
		t = t.parent
	}
	fmt.Printf("t: %#v\n", t)
	return t
}

func (n *Node) treeDelete() {
	var t, c *Node
	if n.left == nil || n.right == nil {
		t = n
	} else {
		t = n.treeSuccessor()
	}

	if t.left != nil {
		c = t.left
	}
	if t.right != nil {
		c = t.right
	}

	if c != nil {
		c.parent = t.parent
	}

	if t.parent == nil {
		n = c
	} else if t == t.parent.left {
		t.parent.left = c
	} else {
		fmt.Println("ok")
		t.parent.right = c
	}
	if t != n {
		n.key = t.key
	}
}

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
		switch sc.Text() {
		case "insert":
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			root.insert(k)
		case "print":
			root.inorder()
			out.WriteString("\n")
			root.preorder()
			out.WriteString("\n")
		case "find":
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			t := root.find(k)
			if t != nil {
				out.WriteString("yes\n")
			} else {
				out.WriteString("no\n")
			}
		case "delete":
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			t := root.find(k)
			if t != nil {
				t.treeDelete()
			}
		}
	}

	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
