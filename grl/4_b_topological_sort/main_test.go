package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "6 6\n0 1\n1 2\n3 1\n3 4\n4 5\n5 2\n",
			want:  "0\n3\n1\n4\n5\n2",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
