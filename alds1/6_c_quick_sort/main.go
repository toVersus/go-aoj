package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// SENTINEL is used for merge sort by adding the last index
var SENTINEL = 2000000000

type Card struct {
	suit  string
	value int
}

func merge(A []Card, n, left, mid, right int) {
	n1 := mid - left
	L := make([]Card, n1+1)
	n2 := right - mid
	R := make([]Card, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}
	L[n1].value, R[n2].value = SENTINEL, SENTINEL

	i, j := 0, 0
	for k := left; k < right; k++ {
		if L[i].value <= R[j].value {
			A[k] = L[i]
			i++
			continue
		}
		A[k] = R[j]
		j++
	}
}

func mergeSort(A []Card, n, left, right int) {
	mid := 0
	if left+1 < right {
		mid = (left + right) / 2
		mergeSort(A, n, left, mid)
		mergeSort(A, n, mid, right)
		merge(A, n, left, mid, right)
	}
}

func partition(A []Card, n, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j].value <= x.value {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func quickSort(A []Card, n, p, r int) {
	q := 0
	if p < r {
		q = partition(A, n, p, r)
		quickSort(A, n, p, q-1)
		quickSort(A, n, q+1, r)
	}
}

func answer(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	A, B := make([]Card, n), make([]Card, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		ss := strings.Split(sc.Text(), " ")
		A[i].suit = ss[0]
		A[i].value, _ = strconv.Atoi(ss[1])
		B[i].suit, B[i].value = A[i].suit, A[i].value
	}

	mergeSort(A, n, 0, n)
	quickSort(B, n, 0, n-1)

	var buf bytes.Buffer
	if reflect.DeepEqual(A, B) {
		buf.WriteString("Stable\n")
	} else {
		buf.WriteString("Not stable\n")
	}

	for i := 0; i < n; i++ {
		buf.WriteString(B[i].suit)
		buf.WriteString(" ")
		buf.WriteString(strconv.Itoa(B[i].value))
		buf.WriteString("\n")
	}
	return strings.TrimRight(buf.String(), "\n")
}

func main() {
	fmt.Println(answer(os.Stdin))
}
