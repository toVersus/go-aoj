package main

import (
	"reflect"
	"strings"
	"testing"
)

var partitionTests = []struct {
	name string
	text string
	want string
}{
	{
		name: "should do the partition selecting the last element as a pivot element",
		text: "12\n13 19 9 5 12 8 7 4 21 2 6 11\n",
		want: "9 5 8 7 4 2 6 [11] 21 13 19 12",
	},
	{
		name: "should do the partition selecting the last element as a pivot element from large input",
		text: "100\n0 5933 7743 6262 1529 9700 5508 9752 7256 7256 8119 9711 8351 2843 8705 4108 6393 8330 7366 2169 7932 6917 9847 7972 2868 6980 6223 8549 7592 8164 6169 6551 427 6190 2624 635 4920 7944 9310 9862 8484 7363 3301 4710 7236 1876 431 9929 397 675 491 4190 9344 8134 8425 1629 4030 2727 9126 1743 7334 4104 4760 9749 5620 256 932 5572 1613 2490 4509 7119 7405 5695 4049 327 6719 3497 6824 2596 8649 7356 4184 3093 2245 1007 9306 6509 2754 352 5665 6783 8738 801 9690 1330 7337 8195 6656 5963",
		want: "0 5933 1529 5508 2843 4108 2169 2868 427 2624 635 4920 3301 4710 1876 431 397 675 491 4190 1629 4030 2727 1743 4104 4760 5620 256 932 5572 1613 2490 4509 5695 4049 327 3497 2596 4184 3093 2245 1007 2754 352 5665 801 1330 [5963] 6393 8330 7366 6262 9344 8134 8425 7932 6917 9847 9126 7972 7334 9752 6980 9749 6223 8549 7592 8164 6169 6551 7256 7119 7405 6190 7256 8119 6719 9711 6824 7944 8649 7356 9310 9862 8484 7363 9306 6509 8351 7743 7236 6783 8738 8705 9690 9700 7337 8195 6656 9929",
	},
}

func TestPartition(t *testing.T) {
	for _, testcase := range partitionTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}
	}
}

var result string

func BenchmarkPartition(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		for _, testcase := range partitionTests {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			str = answer(r)
		}
	}
	result = str
}
