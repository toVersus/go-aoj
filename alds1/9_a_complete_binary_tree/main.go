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

func parent(i int) int { return i / 2 }
func left(i int) int   { return 2 * i }
func right(i int) int  { return 2*i + 1 }

func answer(reader io.Reader) string {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	H, _ := strconv.Atoi(sc.Text())

	A := make([]int, H+1)
	for i := 1; i <= H; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		A[i] = k
	}

	var out bytes.Buffer
	for i := 1; i <= H; i++ {
		out.WriteString(fmt.Sprintf("node %d: key = %d, ", i, A[i]))
		if parent(i) >= 1 {
			out.WriteString(fmt.Sprintf("parent key = %d, ", A[parent(i)]))
		}
		if left(i) <= H {
			out.WriteString(fmt.Sprintf("left key = %d, ", A[left(i)]))
		}
		if right(i) <= H {
			out.WriteString(fmt.Sprintf("right key = %d, ", A[right(i)]))
		}
		out.WriteString("\n")
	}

	return strings.TrimRight(out.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
