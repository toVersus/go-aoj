package main

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var insertionSortTests = []struct {
	name string
	file string
	text string
	want []int
}{
	{
		name: "should be sorted gettint input from large file",
		file: "in5.txt",
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

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		f, err = os.Open(testcase.file)
		if err != nil {
			t.Errorf("could not open a file: %s\n  %s", testcase.file, err)
		}

		sc := bufio.NewScanner(f)
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
		f.Close()

		if result := insertionSort(A, N); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for _, testcase := range insertionSortTests {
		f, err := os.Create(testcase.file)
		if err != nil {
			b.Errorf("could not create a file: %s\n  %s", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range insertionSortTests {
			f, _ := os.Open(testcase.file)
			sc := bufio.NewScanner(f)
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
			f.Close()

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

	for _, testcase := range insertionSortTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
