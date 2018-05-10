package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := 0
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	A := make([]int, n)
	for i := 0; i < n; i++ {
		if !sc.Scan() {
			break
		}
		A[i], _ = strconv.Atoi(sc.Text())
	}

	G, count := []int{}, 0
	A, G, count = shellSort(A, n)
	fmt.Println(len(G))
	fmt.Println(String(G))
	fmt.Println(count)
	fmt.Println(Stringln(A))
}

func insertionSort(A []int, n, g int) ([]int, int) {
	cnt := 0
	for i := g; i < n; i++ {
		v := A[i]
		j := i - g
		for j >= 0 && A[j] > v {
			A[j+g] = A[j]
			j = j - g
			cnt++
		}
		A[j+g] = v
	}
	return A, cnt
}

func shellSort(A []int, n int) ([]int, []int, int) {
	cnt, tmp := 0, 0
	var G []int

	for h := 1; h <= n; {
		G = append(G, 0)
		copy(G[1:], G[0:])
		G[0] = h
		h = 3*h + 1
	}

	for i := 0; i < len(G); i++ {
		A, tmp = insertionSort(A, n, G[i])
		cnt += tmp
	}
	return A, G, cnt
}

func String(A []int) string {
	return strings.Trim(fmt.Sprint(A), "[]")
}

func Stringln(A []int) string {
	numseq := strings.Trim(fmt.Sprint(A), "[]")
	return strings.Replace(numseq, " ", "\n", -1)
}
