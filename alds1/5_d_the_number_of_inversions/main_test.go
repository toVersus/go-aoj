package main

import (
	"reflect"
	"strings"
	"testing"
)

var numInversionTests = []struct {
	name string
	text string
	want string
}{
	{
		name: "should return the number of inversion from simple input",
		text: "5\n3 5 2 1 4",
		want: "6",
	},
	{
		name: "should return the number of inversion from large input",
		text: "100\n1 72 68 4 44 12 7 3 9 32 11 6 83 14 15 5 65 18 19 20 21 22 23 24 25 26 27 28 29 90 35 89 33 34 61 36 8 59 39 40 99 49 43 55 45 64 47 48 42 51 10 97 53 41 30 56 57 58 38 60 31 17 63 46 62 96 67 37 69 70 71 93 73 74 95 76 77 13 98 80 81 82 78 84 85 86 87 88 94 16 91 92 2 66 75 50 52 79 54 100\n",
		want: "1266",
	},
	{
		name: "should return zero due to only one input number",
		text: "1\n1",
		want: "0",
	},
}

func TestNumInversion(t *testing.T) {
	for _, testcase := range numInversionTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}
	}
}

var result string

func BenchmarkMergeSort(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		for _, testcase := range numInversionTests {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			str = answer(r)
		}
	}
	result = str
}
