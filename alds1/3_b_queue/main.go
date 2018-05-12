package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ss := strings.Fields(sc.Text())
	n, _ := strconv.Atoi(ss[0])
	q, _ := strconv.Atoi(ss[1])

	que := newQueue()
	for i := 0; i < n; i++ {
		sc.Scan()
		ss := strings.Fields(sc.Text())
		cputime, _ := strconv.Atoi(ss[1])
		que.enqueue(&process{name: ss[0], time: cputime})
	}
	fmt.Print(que.consume(q))
}

type queue struct {
	head, tail *node
	len        int
}

type node struct {
	proc *process
	next *node
}

type process struct {
	name string
	time int
}

func newQueue() *queue {
	return &queue{head: nil, tail: nil, len: 0}
}

func (q *queue) dequeue() *process {
	if q.len == 0 {
		return nil
	}
	headPtr := q.head
	q.head = headPtr.next
	if q.len == 1 {
		q.head, q.tail = nil, nil
	}
	q.len--
	return headPtr.proc
}

func (q *queue) enqueue(proc *process) {
	n := &node{proc: proc, next: nil}
	if q.len == 0 {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.len++
}

func (q *queue) peek() *process {
	if q.len == 0 {
		return nil
	}
	return q.head.proc
}

func (q *queue) consume(qtm int) string {
	timer, used := 0, 0
	result := ""
	for q.head != nil {
		u := q.dequeue()
		used = min(qtm, u.time)
		u.time -= used
		timer += used
		if u.time > 0 {
			q.enqueue(u)
			continue
		}
		result += u.name + " " + strconv.Itoa(timer) + "\n"
	}
	return strings.TrimRight(result, "\n")
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
