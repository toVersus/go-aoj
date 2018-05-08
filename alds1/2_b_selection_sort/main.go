package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	N := 0
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}
	A := make([]int, N)
	if sc.Scan() {
		for i, val := range strings.Fields(sc.Text()) {
			A[i], _ = strconv.Atoi(val)
		}
	}

	result, count := selectionSort(A, N)
	fmt.Println(String(result))
	fmt.Println(count)
}

func selectionSort(A []int, N int) ([]int, int) {
	swapCount := 0
	for i := 0; i < N-1; i++ {
		minj := i
		for j := i; j < N; j++ {
			if A[j] < A[minj] {
				minj = j
			}
		}
		if i == minj {
			continue
		}
		A[i], A[minj] = A[minj], A[i]
		swapCount++
	}

	return A, swapCount
}

func String(A []int) string {
	return strings.Trim(fmt.Sprint(A), "[]")
}
