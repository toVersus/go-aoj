package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSingleSourceShortestPath1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "5\n0 3 2 3 3 1 1 2\n1 2 0 2 3 4\n2 3 0 3 3 1 4 1\n3 4 2 1 0 1 1 4 4 3\n4 2 2 1 3 3\n",
			want:  "0 0\n1 2\n2 2\n3 1\n4 3",
		},
		{
			input: "9\n0 2 1 1 3 13\n1 3 0 1 2 1 4 11\n2 2 5 1 1 1\n3 3 4 1 0 13 6 1\n4 4 1 11 5 1 3 1 7 4\n5 3 2 1 8 7 4 1\n6 2 3 1 7 1\n7 3 4 4 6 1 8 1\n8 2 5 7 7 1\n",
			want:  "0 0\n1 1\n2 2\n3 5\n4 4\n5 3\n6 6\n7 7\n8 8",
		},
		{
			input: "10\n0 2 1 1 2 3\n1 2 3 3 2 1\n2 2 3 1 6 10\n3 1 4 1\n4 1 7 1\n5 2 3 1 6 1\n6 2 3 1 8 1\n7 2 5 1 8 4\n8 1 9 100000\n9 0\n",
			want:  "0 0\n1 1\n2 2\n3 3\n4 4\n5 6\n6 7\n7 5\n8 8\n9 100008",
		},
		{
			input: "20\n0 3 1 100 2 2 4 5\n1 3 0 100 2 3 3 1\n2 5 1 3 0 2 4 10 5 12 3 1\n3 2 2 1 1 1\n4 3 0 5 2 10 5 2\n5 4 4 2 2 12 6 19 7 3\n6 4 5 19 11 4 10 3 7 17\n7 5 5 3 6 17 10 1000 9 2 8 5\n8 2 7 5 9 5\n9 4 8 5 7 2 10 1 13 50\n10 5 9 1 7 1000 6 3 11 8 12 100\n11 3 12 7 10 8 6 4\n12 5 10 100 11 7 15 2 14 5 13 1\n13 2 12 1 9 50\n14 4 15 2 12 5 18 3 16 3\n15 2 12 2 14 2\n16 2 14 3 17 5\n17 3 18 1 19 1 16 5\n18 3 14 3 17 1 19 3\n19 2 18 3 17 1\n",
			want:  "0 0\n1 4\n2 2\n3 3\n4 5\n5 7\n6 16\n7 10\n8 15\n9 12\n10 13\n11 20\n12 27\n13 28\n14 31\n15 29\n16 34\n17 35\n18 34\n19 36",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
