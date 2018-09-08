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

const N = 1001

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lcs(x, y string) int {
	var c [N][N]int
	m := len(x)
	n := len(y)
	var maxl int

	x = " " + x
	y = " " + y

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i] == y[j] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = max(c[i-1][j], c[i][j-1])
			}
			maxl = max(maxl, c[i][j])
		}
	}
	return maxl
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	var out bytes.Buffer
	for i := 0; i < q; i++ {
		sc.Scan()
		s1 := sc.Text()
		sc.Scan()
		s2 := sc.Text()
		out.WriteString(fmt.Sprintf("%d\n", lcs(s1, s2)))
	}
	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
