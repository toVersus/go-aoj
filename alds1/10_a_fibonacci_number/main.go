package main

import (
	"fmt"
)

func answer(n int) string {
	f := []int{1, 1}
	for i := 2; i <= n; i++ {
		f = append(f, f[i-1]+f[i-2])
	}

	return fmt.Sprintf("%d", f[n])
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(answer(n))
}
