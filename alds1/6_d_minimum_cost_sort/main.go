package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const VMAX = 100000

func main() {
	fmt.Printf("%d\n", answer(os.Stdin))
}

func answer(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	A := make([]int, n)
	s := VMAX
	for i := 0; i < n; i++ {
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
		s = min(s, A[i])
	}
	return minCostSort(A, n, s)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCostSort(A []int, n, s int) int {
	ans := 0
	B, T := make([]int, n), make([]int, VMAX)
	V := make([]bool, n) // all elements are "false".
	copy(B, A)
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})
	for i := 0; i < n; i++ {
		T[B[i]] = i
	}
	for i := 0; i < n; i++ {
		if V[i] {
			continue
		}
		cur := i
		// S is sum of elements.
		S := 0
		// m is minimum value of elements.
		m := VMAX
		// an is answer in cycle.
		an := 0
		for {
			V[cur] = true
			an++
			v := A[cur]
			m = min(m, v) // minimum element.
			S += v        // Sum of elements.
			cur = T[v]
			if V[cur] {
				break
			}
		}
		ans += min(S+(an-2)*m, m+S+(an+1)*s)
	}
	return ans
}
