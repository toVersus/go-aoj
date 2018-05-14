package main

import (
	"reflect"
	"strings"
	"testing"
)

var binarySearchTests = []struct {
	name string
	text string
	want int
}{
	{
		name: "should return correct number of union of slice",
		text: "15\n0 0 1 1 2 2 3 3 3 5 6 6 8 9 9\n10\n8 4 6 5 0 2 1 7 9 3\n",
		want: 8,
	},
	{
		name: "no intersection between two slices",
		text: "20\n0 0 0 1 1 2 2 3 3 3 3 5 6 6 6 8 8 9 9 9\n10\n19 16 12 14 10 15 11 13 17 18\n",
		want: 0,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, testcase := range binarySearchTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
	}
}
