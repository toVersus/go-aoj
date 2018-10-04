package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestRangeSearch(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "6\n2 1\n2 2\n4 2\n6 2\n3 3\n5 4\n2\n2 4 0 4\n4 10 2 5\n",
			want:  "0\n1\n2\n4\n\n2\n3\n5\n\n",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
