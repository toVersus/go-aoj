package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDisjointSet(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "5 12\n0 1 4\n0 2 3\n1 1 2\n1 3 4\n1 1 4\n1 3 2\n0 1 3\n1 2 4\n1 3 0\n0 0 4\n1 0 2\n1 3 0\n",
			want:  "0\n0\n1\n1\n1\n0\n1\n1",
		},
		{
			input: "10 18\n0 0 1\n1 5 9\n0 1 2\n1 5 2\n0 2 3\n1 6 3\n0 3 4\n1 2 9\n0 4 5\n1 2 1\n0 5 6\n1 7 5\n0 6 7\n1 8 6\n0 7 8\n1 4 6\n0 8 9\n1 3 1\n",
			want:  "0\n0\n0\n0\n1\n0\n0\n1\n1",
		},
		{
			input: "16 20\n1 8 0\n1 1 13\n0 10 14\n0 5 4\n0 10 11\n1 1 5\n0 6 2\n1 0 9\n1 14 15\n0 13 9\n0 8 12\n0 1 2\n0 1 5\n0 8 4\n1 6 8\n0 2 3\n0 7 6\n1 11 10\n0 13 12\n0 8 9\n",
			want:  "0\n0\n0\n0\n0\n1\n1",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
