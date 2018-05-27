package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// VMAX is the maximum value of element in progression.
const VMAX = 10000

func main() {
	fmt.Println(answer(os.Stdin))
}

func answer(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// A[0], B[0] and C[0] won't be used due to handling the inconsistency
	// between the size of slice and its number of element.
	A, B := make([]int, n+1), make([]int, n+1)
	C := make([]int, VMAX+1)
	for i := 0; i < n; i++ {
		sc.Scan()
		A[i+1], _ = strconv.Atoi(sc.Text())
		C[A[i+1]]++
	}

	countingSort(A, B, C, n)
	return String(B)
}

func countingSort(A, B, C []int, n int) {
	for i := 1; i <= VMAX; i++ {
		C[i] += C[i-1]
	}

	for j := 1; j <= n; j++ {
		B[C[A[j]]] = A[j]
		C[A[j]]--
	}
}

func String(B []int) string {
	// trim junk element B[0] when stringifying the entire elements of progression.
	return strings.Trim(fmt.Sprint(B[1:])[1:], "[]")
}
