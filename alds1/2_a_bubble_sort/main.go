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
	result, count := bubbleSort(A, N)
	fmt.Println(String(result))
	fmt.Println(count)
}

func bubbleSort(A []int, N int) ([]int, int) {
	flag := true
	count := 0
	for flag {
		flag = false
		for j := N - 1; j > 0; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
				count++
				flag = true
			}
		}
	}
	return A, count
}

func String(A []int) string {
	return strings.Trim(fmt.Sprint(A), "[]")
}
