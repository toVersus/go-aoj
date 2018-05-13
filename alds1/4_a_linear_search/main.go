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
	count := answer(os.Stdin)
	fmt.Println(count)
}

func answer(r io.Reader) int {
	sc := bufio.NewScanner(r)

	// catch a first input line representing the length of slices to be passed in second input.
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// catch a second input line representing the elements of slices.
	sc.Scan()
	A := make([]int, n+1)
	sa := strings.Fields(sc.Text())
	for i := 0; i < n; i++ {
		A[i], _ = strconv.Atoi(sa[i])
	}

	// catch a third input line representing the length of slices to be passed in fourth input.
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	// catch a fourth input line representing the elements of slices.
	sc.Scan()
	count := 0
	st := strings.Fields(sc.Text())
	for j := 0; j < q; j++ {
		key, _ := strconv.Atoi(st[j])
		if !search(A, n, key) {
			continue
		}
		count++
	}
	return count
}

func search(A []int, n, key int) bool {
	// add sentinel to the end of the slice to reduce the nest of conditions.
	A = append(A, key)
	i := 0
	for A[i] != key {
		i++
	}
	return i != n+1
}
