package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	N := 0
	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}
	A := make([]int, N)
	if sc.Scan() {
		for i, s := range strings.Fields(sc.Text()) {
			A[i], _ = strconv.Atoi(s)
		}
	}
	fmt.Println(String(A))
	insertionSort(A, N)
}

func insertionSort(A []int, N int) []int {
	for i := 1; i < N; i++ {
		v := A[i]
		j := i - 1
		for j >= 0 && A[j] > v {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
		fmt.Println(String(A))
	}
	return A
}

func String(A []int) string {
	return strings.Trim(fmt.Sprint(A), "[]")
}
