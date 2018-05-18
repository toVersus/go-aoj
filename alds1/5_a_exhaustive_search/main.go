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

var (
	n int
	A []int
)

func main() {
	fmt.Println(answer(os.Stdin))
}

func solve(i, m int) bool {
	if m == 0 {
		return true
	}
	if i >= n {
		return false
	}
	return solve(i+1, m) || solve(i+1, m-A[i])
}

func answer(r io.Reader) string {
	sc := bufio.NewScanner(r)
	var buf bytes.Buffer
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	A = make([]int, n)
	ss := strings.Fields(sc.Text())
	for i := 0; i < n; i++ {
		A[i], _ = strconv.Atoi(ss[i])
	}
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	ss = strings.Fields(sc.Text())
	for j := 0; j < m; j++ {
		M, _ := strconv.Atoi(ss[j])
		if solve(0, M) {
			buf.WriteString("yes\n")
			continue
		}
		buf.WriteString("no\n")
	}
	return strings.TrimRight(buf.String(), "\n")
}
