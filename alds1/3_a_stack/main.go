package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The basic implementation of stack is inspired by the following code:
// https://github.com/golang-collections/collections/blob/master/stack/stack.go

func main() {
	op := []string{}
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		op = strings.Fields(sc.Text())
	}
	fmt.Println(compute(op))
}

func compute(op []string) int {
	r, l := -1, -1
	st := newStack()
	for i := 0; i < len(op); i++ {
		switch op[i] {
		case "+":
			r, l = st.pop(), st.pop()
			st.push(l + r)
		case "-":
			r, l = st.pop(), st.pop()
			st.push(l - r)
		case "*":
			r, l = st.pop(), st.pop()
			st.push(l * r)
		default:
			num, err := strconv.Atoi(op[i])
			if err != nil {
				fmt.Println(err)
			}
			st.push(num)
		}
	}
	return st.peek()
}

type stack struct {
	top *node
	len int
}

type node struct {
	value int
	prev  *node
}

func newStack() *stack {
	return &stack{nil, 0}
}

func (s *stack) peek() int {
	if s.len == 0 {
		return -1
	}
	return s.top.value
}

func (s *stack) pop() int {
	if s.len == 0 {
		return -1
	}
	n := s.top
	s.top = n.prev
	s.len--
	return n.value
}

func (s *stack) push(num int) {
	s.top = &node{value: num, prev: s.top}
	s.len++
}
