package main

import (
	"reflect"
	"strings"
	"testing"
)

var quickSortTests = []struct {
	name string
	text string
	want string
}{
	{
		name: "should return Cards sorted by their value in non-stable order",
		text: "6\nD 3\nH 2\nD 1\nS 3\nD 2\nC 1\n",
		want: "Not stable\nD 1\nC 1\nD 2\nH 2\nD 3\nS 3",
	},
	{
		name: "should return original Cards without sorting, so they keep stability",
		text: "2\nS 1\nH 1\n",
		want: "Stable\nS 1\nH 1",
	},
	{
		name: "should return Cards sorted by their value in non-stable order from large input",
		text: "40\nH 20\nD 20\nC 20\nS 20\nH 30\nD 30\nH 10\nD 10\nC 10\nS 10\nC 30\nS 30\nH 40\nH 3\nD 3\nC 3\nS 3\nH 4\nD 4\nD 40\nC 40\nS 40\nC 50\nS 50\nH 1\nD 1\nC 1\nH 50\nD 50\nS 1\nH 2\nD 2\nC 2\nS 2\nC 4\nS 4\nH 5\nD 5\nC 5\nS 5\n",
		want: "Not stable\nH 1\nD 1\nC 1\nS 1\nH 2\nD 2\nC 2\nS 2\nC 3\nS 3\nH 3\nD 3\nH 4\nD 4\nC 4\nS 4\nH 5\nD 5\nC 5\nS 5\nH 10\nD 10\nC 10\nS 10\nH 20\nD 20\nC 20\nS 20\nC 30\nS 30\nH 30\nD 30\nC 40\nS 40\nH 40\nD 40\nD 50\nS 50\nC 50\nH 50",
	},
}

func TestQuickSort(t *testing.T) {
	for _, testcase := range quickSortTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}
	}
}
