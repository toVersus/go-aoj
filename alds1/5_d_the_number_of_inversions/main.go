package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	SENTINEL = 2000000000
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
	return strconv.Itoa(mergeSort(A, n, 0, n))
}

func merge(A []int, n, left, mid, right int) int {
	n1 := mid - left
	n2 := right - mid
	L, R := make([]int, n1+1), make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}
	L[n1], R[n2] = SENTINEL, SENTINEL

	cnt := 0
	i, j := 0, 0
	for k := left; k < right; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
			continue
		}
		A[k] = R[j]
		j++
		cnt += n1 - i
	}
	return cnt
}

func mergeSort(A []int, n, left, right int) int {
	if left+1 < right {
		mid := (left + right) / 2
		v1 := mergeSort(A, n, left, mid)
		v2 := mergeSort(A, n, mid, right)
		v3 := merge(A, n, left, mid, right)
		return v1 + v2 + v3
	}
	return 0
}
