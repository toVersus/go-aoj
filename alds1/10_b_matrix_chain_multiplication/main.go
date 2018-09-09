package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const N = 100

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var p [N]int
	var m [N + 1][N + 1]int
	for i := 1; i <= n; i++ {
		sc.Scan()
		p[i-1], _ = strconv.Atoi(sc.Text())
		sc.Scan()
		p[i], _ = strconv.Atoi(sc.Text())
	}

	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i][j] = (1 << 21)
			for k := i; k <= j-1; k++ {
				m[i][j] = min(m[i][j], m[i][k]+m[k+1][j]+p[i-1]*p[k]*p[j])
			}
		}
	}

	return fmt.Sprintf("%d", m[1][n])
}

func main() {
	fmt.Println(answer(os.Stdin))
}
