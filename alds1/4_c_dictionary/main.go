package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// M is large primary number
const M = 1046527

func main() {
	fmt.Println(answer(os.Stdin))
}

func answer(r io.Reader) string {
	H := make([]string, M)
	sc := bufio.NewScanner(r)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	buf := bytes.Buffer{}
	for i := 0; i < n; i++ {
		sc.Scan()
		ss := strings.Fields(sc.Text())
		switch ss[0] {
		case "insert":
			H = insert(H, ss[1])
		case "find":
			if find(H, ss[1]) {
				buf.WriteString("yes\n")
			} else {
				buf.WriteString("no\n")
			}
		}
	}
	return strings.TrimRight(buf.String(), "\n")
}

func atoi(char byte) int {
	switch char {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	default:
		return 0
	}
}

func getKey(str string) int {
	sum, p := 0, 1
	for i := 0; i < len(str); i++ {
		sum += p * atoi(str[i])
		p *= 5
	}
	return sum
}

func h1(key int) int {
	return key % M
}

func h2(key int) int {
	return 1 + key%(M-1)
}

func find(H []string, str string) bool {
	key := getKey(str)
	h := 0
	for i := 0; ; i++ {
		h = (h1(key) + i*h2(key)) % M
		if H[h] == str {
			return true
		} else if len(H[h]) == 0 {
			return false
		}
	}
}

func insert(H []string, str string) []string {
	key := getKey(str)
	for i := 0; ; i++ {
		h := (h1(key) + i*h2(key)) % M
		if H[h] == str {
			return H
		} else if len(H[h]) == 0 {
			H[h] = str
			return H
		}
	}
}
