package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestMinimumSpanningTree(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "5\n -1 2 3 1 -1\n 2 -1 -1 4 -1\n 3 -1 -1 1 1\n 1 4 1 -1 3\n -1 -1 1 3 -1\n",
			want:  "5",
		},
		{
			input: "10\n -1 1 3 -1 -1 -1 -1 -1 -1 -1\n 1 -1 1 3 -1 -1 -1 -1 -1 -1\n 3 1 -1 1 -1 -1 10 -1 -1 -1\n -1 3 1 -1 11 1 -1 -1 -1\n -1 -1 -1 1 -1 -1 -1 1 -1 -1\n -1 -1 -1 1 -1 -1 1 1 -1 -1\n -1 -1 10 1 -1 1 -1 -1 1 -1\n -1-1 -1 -1 1 1 -1 -1 4 -1\n -1 -1 -1 -1 -1 -1 1 4 -1 1000\n -1 -1 -1 -1 -1 -1 -1 -1 1000 -1\n",
			want:  "1008",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
