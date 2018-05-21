package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	MAX      = 500000
	SENTINEL = 2000000000
	cnt      = 0
	L, R     = make([]int, MAX+2), make([]int, MAX+2)
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

	mergeSort(A, n, 0, n)

	result := strings.TrimRight(fmt.Sprintf("%+v", A)[1:], "]") + "\n"
	return result + strconv.Itoa(cnt)
}

func merge(A []int, n, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid
	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}
	L[n1], R[n2] = SENTINEL, SENTINEL

	i, j := 0, 0
	for k := left; k < right; k++ {
		cnt++
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
			continue
		}
		A[k] = R[j]
		j++
	}
}

func mergeSort(A []int, n, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(A, n, left, mid)
		mergeSort(A, n, mid, right)
		merge(A, n, left, mid, right)
	}
}
