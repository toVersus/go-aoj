package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(answer(os.Stdin))
}

func answer(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	// catch a first input line representing the length of slices to be passed in second input.
	n := nextInt(sc)

	// catch a second input line representing the elements of slices.
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = nextInt(sc)
	}

	// catch a third input line representing the length of slices to be passed in fourth input.
	q := nextInt(sc)

	// catch a fourth input line representing the elements of slices.
	count := 0
	for j := 0; j < q; j++ {
		key := nextInt(sc)
		if !binarySearch(A, n, key) {
			continue
		}
		count++
	}
	return count
}

func binarySearch(A []int, lenA, key int) bool {
	left, mid, right := 0, 0, lenA
	for left < right {
		mid = (left + right) / 2
		if key == A[mid] {
			return true
		} else if key > A[mid] { // to search second half of the slice
			left = mid + 1
			continue
		}
		// to search first half of the slice
		right = mid
	}
	return false
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
