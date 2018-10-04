package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	NIL = -1
)

var np = 0

type Node struct {
	location int
	p, l, r  int
}

type Point struct {
	id, x, y int
}

func (p *Point) print() string {
	return fmt.Sprintln(p.id)
}

type Points []Point

func (p Points) makekDtree(T []Node, l, r, dep int) int {
	if !(l < r) {
		return NIL
	}
	mid := (l + r) / 2
	t := np
	np++

	if dep%2 == 0 {
		sort.Slice(p, func(i, j int) bool {
			return p[i].x > p[j].x
		})
	} else {
		sort.Slice(p, func(i, j int) bool {
			return p[i].y > p[j].y
		})
	}
	T[t].location = mid
	T[t].l = p.makekDtree(T, l, mid, dep+1)
	T[t].r = p.makekDtree(T, mid+1, r, dep+1)

	return t
}

func (p Points) find(T []Node, v, sx, tx, sy, ty, dep int, ans *Points) {
	x := p[T[v].location].x
	y := p[T[v].location].y

	if sx <= x && x <= tx && sy <= y && y <= ty {
		*ans = append(*ans, p[T[v].location])
	}

	if dep%2 == 0 {
		if T[v].l != NIL && sx <= x {
			p.find(T, T[v].l, sx, tx, sy, ty, dep+1, ans)
		}
		if T[v].r != NIL && x <= tx {
			p.find(T, T[v].r, sx, tx, sy, ty, dep+1, ans)
		}
	} else {
		if T[v].l != NIL && sy <= y {
			p.find(T, T[v].l, sx, tx, sy, ty, dep+1, ans)
		}
		if T[v].r != NIL && y <= ty {
			p.find(T, T[v].r, sx, tx, sy, ty, dep+1, ans)
		}
	}
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	P, T := make(Points, n), make([]Node, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		y, _ := strconv.Atoi(sc.Text())
		P[i] = Point{i, x, y}
		T[i].l, T[i].r, T[i].p = NIL, NIL, NIL
	}

	root := P.makekDtree(T, 0, n, 0)

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	var out bytes.Buffer
	for i := 0; i < q; i++ {
		sc.Scan()
		sx, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		tx, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		sy, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		ty, _ := strconv.Atoi(sc.Text())

		var ans Points
		P.find(T, root, sx, tx, sy, ty, 0, &ans)
		sort.Slice(ans, func(i, j int) bool {
			return ans[i].id < ans[j].id
		})
		for j := 0; j < len(ans); j++ {
			out.WriteString(ans[j].print())
		}
		out.WriteString("\n")
	}

	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
