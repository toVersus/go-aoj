package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(answer(os.Stdin))
}

func answer(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	A := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}
	q := partition(A, 0, n-1)

	pref := strings.TrimRight(fmt.Sprintf("%+v", A[:q])[1:], "]")
	suf := strings.TrimRight(fmt.Sprintf("%+v", A[q+1:])[1:], "]")
	return pref + " [" + strconv.Itoa(A[q]) + "] " + suf
}

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] > x {
			continue
		}
		i++
		A[i], A[j] = A[j], A[i]
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
