package main

import (
	"bufio"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var selectionSortTests = []struct {
	name     string
	text     string
	want     []int
	attempts int
}{
	{
		name: "just print out the single input number without sorting",
		text: `1
100
`,
		want:     []int{100},
		attempts: 0,
	},
	{
		name: "should print out the input numbers without sorting",
		text: `5
1 2 3 4 5
`,
		want:     []int{1, 2, 3, 4, 5},
		attempts: 0,
	},
	{
		name: "should print out the sorted input numbers in maximum number of attempts",
		text: `6
6 5 4 3 2 1
`,
		want:     []int{1, 2, 3, 4, 5, 6},
		attempts: 3,
	},
	{
		name: "should print out the sorted 10 input numbers",
		text: `10
9 5 10 3 6 7 4 1 2 8
`,
		want:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		attempts: 9,
	},
	{
		name: "should print out the sorted bunch of numbers getting from large input",
		text: `100
0 33 43 62 29 0 8 52 56 56 19 11 51 43 5 8 93 30 66 69 32 17 47 72 68 80 23 49 92 64 69 51 27 90 24 35 20 44 10 62 84 63 1 10 36 76 31 29 97 75 91 90 44 34 25 29 30 27 26 43 34 4 60 49 20 56 32 72 13 90 9 19 5 95 49 27 19 97 24 96 49 56 84 93 45 7 6 9 54 52 65 83 38 1 90 30 37 95 56 63
`,
		want:     []int{0, 0, 1, 1, 4, 5, 5, 6, 7, 8, 8, 9, 9, 10, 10, 11, 13, 17, 19, 19, 19, 20, 20, 23, 24, 24, 25, 26, 27, 27, 27, 29, 29, 29, 30, 30, 30, 31, 32, 32, 33, 34, 34, 35, 36, 37, 38, 43, 43, 43, 44, 44, 45, 47, 49, 49, 49, 49, 51, 51, 52, 52, 54, 56, 56, 56, 56, 56, 60, 62, 62, 63, 63, 64, 65, 66, 68, 69, 69, 72, 72, 75, 76, 80, 83, 84, 84, 90, 90, 90, 90, 91, 92, 93, 93, 95, 95, 96, 97, 97},
		attempts: 95,
	},
}

func TestSelectionSort(t *testing.T) {
	for _, testcase := range selectionSortTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)

		N := 0
		sc := bufio.NewScanner(r)
		if sc.Scan() {
			N, _ = strconv.Atoi(sc.Text())
		}
		A := make([]int, N)
		if sc.Scan() {
			for i, val := range strings.Fields(sc.Text()) {
				A[i], _ = strconv.Atoi(val)
			}
		}

		result, counts := selectionSort(A, N)
		if !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
		if !reflect.DeepEqual(counts, testcase.attempts) {
			t.Errorf("result => %#v\n, want => %#v", counts, testcase.attempts)
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range selectionSortTests {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			sc := bufio.NewScanner(r)
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

			selectionSort(A, N)
		}
	}
}
