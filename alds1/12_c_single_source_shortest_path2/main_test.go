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
			input: "9\n0 2 1 1 3 13\n1 3 0 1 2 1 4 11\n2 2 5 1 1 1\n3 3 4 1 0 13 6 1\n4 4 1 11 5 1 3 1 7 4\n5 3 2 1 8 7 4 1\n6 2 3 1 7 1\n7 3 4 4 6 1 8 1\n8 2 5 7 7 1\n",
			want:  "0 0\n1 1\n2 2\n3 5\n4 4\n5 3\n6 6\n7 7\n8 8",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
