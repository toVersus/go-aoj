package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

const th = math.Pi * 60.0 / 180.0

var buf bytes.Buffer

type point struct {
	x float64
	y float64
}

func main() {
	fmt.Print(answer(os.Stdin))
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Scan()
	d, _ := strconv.Atoi(sc.Text())

	var l, r point
	l.x, l.y = 0, 0
	r.x, r.y = 100, 0

	buf.WriteString(fmt.Sprintf("%.8f %.8f\n", l.x, l.y))
	koch(d, l, r)
	buf.WriteString(fmt.Sprintf("%.8f %.8f\n", r.x, r.y))

	return buf.String()
}

func koch(dep int, l, r point) {
	if dep == 0 {
		return
	}
	var s, t, u point

	s.x = (2.0*l.x + 1.0*r.x) / 3.0
	s.y = (2.0*l.y + 1.0*r.y) / 3.0
	t.x = (1.0*l.x + 2.0*r.x) / 3.0
	t.y = (1.0*l.y + 2.0*r.y) / 3.0
	// rotation matrix as the starting point from point "s"
	u.x = s.x + (t.x-s.x)*math.Cos(th) - (t.y-s.y)*math.Sin(th)
	u.y = s.y + (t.x-s.x)*math.Sin(th) + (t.y-s.y)*math.Cos(th)

	koch(dep-1, l, s)
	buf.WriteString(fmt.Sprintf("%.8f %.8f\n", s.x, s.y))
	koch(dep-1, s, u)
	buf.WriteString(fmt.Sprintf("%.8f %.8f\n", u.x, u.y))
	koch(dep-1, u, t)
	buf.WriteString(fmt.Sprintf("%.8f %.8f\n", t.x, t.y))
	koch(dep-1, t, r)
}
