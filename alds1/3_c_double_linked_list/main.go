package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	prev, next *node
	key        int
}

// dummy node represents sentinel node.
var dummy = &node{}

func initNode() {
	dummy.next = dummy
	dummy.prev = dummy
}

// insert inserts the new node at the top position of node.
func insert(key int) {
	new := &node{key: key}
	new.next = dummy.next
	dummy.next.prev = new
	dummy.next = new
	new.prev = dummy
}

// search finds and returns the first node containing the specified key.
func search(key int) *node {
	current := dummy.next
	for current.key != key {
		if current == dummy {
			return nil
		}
		current = current.next
	}
	return current
}

// deleteNode deletes specified node.
func deleteNode(n *node) {
	if n == nil {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n = nil
}

func deleteFirst() {
	deleteNode(dummy.next)
}

func deleteLast() {
	deleteNode(dummy.prev)
}

func stringList() string {
	buf := bytes.Buffer{}
	current := dummy.next
	for {
		buf.WriteString(strconv.Itoa(current.key))
		current = current.next
		if current == dummy {
			break
		}
		buf.WriteString(" ")
	}
	return buf.String()
}

func main() {
	if err := runInst(os.Stdin); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(stringList())
}

// runInst runs the bunch of instructions parsing from the input.
func runInst(fd *os.File) error {
	n := 0
	sc := bufio.NewScanner(fd)
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	initNode()
	nodeKey := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		ss := strings.Fields(sc.Text())
		switch ss[0] {
		case "insert":
			nodeKey, _ = strconv.Atoi(ss[1])
			insert(nodeKey)
		case "delete":
			nodeKey, _ = strconv.Atoi(ss[1])
			if target := search(nodeKey); target != nil {
				deleteNode(target)
			}
		case "deleteFirst":
			deleteFirst()
		case "deleteLast":
			deleteLast()
		default:
			return fmt.Errorf("specify the command form the following list: \"insert <key>\", \"delete <key>\", \"deleteFirst\", \"deleteLast\"")
		}
	}
	return nil
}
