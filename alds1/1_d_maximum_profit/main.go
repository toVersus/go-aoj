package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(getMaxProfitByFscan(os.Stdin))
}

func getMaxProfitByFscan(fd *os.File) int {
	n, val, minv, maxv := 0, 0, 0, -2147483648
	fmt.Fscan(fd, &n)
	fmt.Fscan(fd, &minv)
	for i := 1; i < n; i++ {
		fmt.Fscan(fd, &val)
		if maxv < val-minv {
			maxv = val - minv
		}
		if minv > val {
			minv = val
		}
	}
	return maxv
}

func getMaxProfitByScannerAtoi(fd *os.File) int {
	n, val, minv, maxv := 0, 0, 0, -2147483648
	sc := bufio.NewScanner(fd)
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	if sc.Scan() {
		minv, _ = strconv.Atoi(sc.Text())
	}
	for i := 1; i < n; i++ {
		sc.Scan()
		val, _ = strconv.Atoi(sc.Text())
		if maxv < val-minv {
			maxv = val - minv
		}
		if minv > val {
			minv = val
		}
	}
	return maxv
}