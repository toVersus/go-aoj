package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "3\n1 1 2\n2 1 3\n3 1 1\n",
			want:  "1 1 6\n2 2 5\n3 3 4",
		},
		{
			input: "4\n1 1 2\n2 1 4\n3 0\n4 1 3\n",
			want:  "1 1 8\n2 2 7\n3 4 5\n4 3 6",
		},
		{
			input: "6\n1 2 2 3\n2 2 3 4\n3 1 5\n4 1 6\n5 1 6\n6 0\n",
			want:  "1 1 12\n2 2 11\n3 3 8\n4 9 10\n5 4 7\n6 5 6",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
