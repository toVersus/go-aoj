package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGraph(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "4\n1 2 2 4\n2 1 4\n3 0\n4 1 3\n",
			want:  "0 1 0 1\n0 0 0 1\n0 0 0 0\n0 0 1 0",
		},
		{
			input: "6\n1 2 2 4\n2 1 5\n3 2 5 6\n4 0\n5 1 4\n6 1 6\n",
			want:  "0 1 0 1 0 0\n0 0 0 0 1 0\n0 0 0 0 1 1\n0 0 0 0 0 0\n0 0 0 1 0 0\n0 0 0 0 0 1",
		},
		{
			input: "10\n1 3 2 3 4\n2 0\n3 1 5\n4 1 10\n5 5 6 7 8 9 10\n6 1 2\n7 0\n8 1 7\n9 1 10\n10 1 3\n",
			want:  "0 1 1 1 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0\n0 0 0 0 1 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 1\n0 0 0 0 0 1 1 1 1 1\n0 1 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 1 0 0 0\n0 0 0 0 0 0 0 0 0 1\n0 0 1 0 0 0 0 0 0 0",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
