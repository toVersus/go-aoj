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

const N = 100

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var M [N][N]int
	for i := 0; i < n; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		u--
		for j := 0; j < k; j++ {
			sc.Scan()
			v, _ := strconv.Atoi(sc.Text())
			v--
			M[u][v] = 1
		}
	}

	var out bytes.Buffer
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != 0 {
				out.WriteString(" ")
			}
			out.WriteString(fmt.Sprintf("%d", M[i][j]))
		}
		out.WriteString("\n")
	}

	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
