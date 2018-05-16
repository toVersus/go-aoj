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
	fmt.Println(answer(os.Stdin))
}

func answer(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	ss := strings.Fields(sc.Text())
	n, _ := strconv.Atoi(ss[0])
	k, _ := strconv.Atoi(ss[1])

	T := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		T[i], _ = strconv.Atoi(sc.Text())
	}

	left, mid := 0, 0
	right := 100000 * 100000
	for right-left > 1 {
		mid = (left + right) / 2
		v := allocateCapacity(T, mid, k, n)
		if v >= n {
			right = mid
		} else {
			left = mid
		}
	}
	return strconv.Itoa(right)
}

// allocateCapacity assigns the capacity per unit
// form the max capacity "P", max number of containers "k", max number of balls "n"
func allocateCapacity(T []int, P, k, n int) int {
	i := 0
	for j := 0; j < k; j++ {
		s := 0
		for s+T[i] <= P {
			s += T[i]
			i++
			if i == n {
				return n
			}
		}
	}
	return i
}
