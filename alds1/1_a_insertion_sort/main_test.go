package main

import (
	"bufio"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var insertionSortTests = []struct {
	name string
	text string
	want []int
}{
	{
		name: "should be sorted gettint input from large file",
		text: `100
0 933 743 262 529 700 508 752 256 256 119 711 351 843 705 108 393 330 366 169 932 917 847 972 868 980 223 549 592 164 169 551 427 190 624 635 920 944 310 862 484 363 301 710 236 876 431 929 397 675 491 190 344 134 425 629 30 727 126 743 334 104 760 749 620 256 932 572 613 490 509 119 405 695 49 327 719 497 824 596 649 356 184 93 245 7 306 509 754 352 665 783 738 801 690 330 337 195 656 963
`,
		want: []int{0, 7, 30, 49, 93, 104, 108, 119, 119, 126, 134, 164, 169, 169, 184, 190, 190, 195, 223, 236, 245, 256, 256, 256, 262, 301, 306, 310, 327, 330, 330, 334, 337, 344, 351, 352, 356, 363, 366, 393, 397, 405, 425, 427, 431, 484, 490, 491, 497, 508, 509, 509, 529, 549, 551, 572, 592, 596, 613, 620, 624, 629, 635, 649, 656, 665, 675, 690, 695, 700, 705, 710, 711, 719, 727, 738, 743, 743, 749, 752, 754, 760, 783, 801, 824, 843, 847, 862, 868, 876, 917, 920, 929, 932, 932,
			933, 944, 963, 972, 980},
	},
}

func TestInsertionSort(t *testing.T) {
	for _, testcase := range insertionSortTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
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

		if result := insertionSort(A, N); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range insertionSortTests {
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

			// insertion sort
			for i := 1; i < N; i++ {
				v := A[i]
				j := i - 1
				for j >= 0 && A[j] > v {
					A[j+1] = A[j]
					j--
				}
				A[j+1] = v
			}
		}
	}
}
