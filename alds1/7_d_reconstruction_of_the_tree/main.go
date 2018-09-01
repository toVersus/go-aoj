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

var pos int

type Order struct {
	pre, in, post []int
}

func (o *Order) rec(l, r int) {
	if l >= r {
		return
	}
	root := o.pre[pos]
	pos++

	var m int
	for i, elem := range o.in {
		if elem == root {
			m = i
			break
		}
	}

	o.rec(l, m)
	o.rec(m+1, r)
	o.post = append(o.post, root)
}

func (o *Order) solve() string {
	var out bytes.Buffer

	o.rec(0, len(o.pre))
	for i := 0; i < len(o.pre); i++ {
		out.WriteString(fmt.Sprintf("%d ", o.post[i]))
	}

	return strings.TrimRight(out.String(), " ")
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var order Order
	for i := 0; i < n; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		order.pre = append(order.pre, k)
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		order.in = append(order.in, k)
	}

	return order.solve()
}

func main() {
	fmt.Println(answer(os.Stdin))
}
