package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestMaximumHeap(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "10\n4 1 3 2 16 9 10 14 8 7",
			want:  " 16 14 10 8 7 9 3 2 4 1",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
