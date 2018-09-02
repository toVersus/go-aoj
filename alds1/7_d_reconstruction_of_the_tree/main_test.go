package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReconstructPostOrder(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "5\n1 2 3 4 5\n3 2 4 1 5",
			want:  "3 4 2 5 1",
		},
		{
			input: "4\n1 2 3 4\n4 3 2 1",
			want:  "4 3 2 1",
		},
		{
			input: "1\n1\n1",
			want:  "1",
		},
		{
			input: "30\n1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 24 27 29 30 25 28 23 26\n6 5 7 8 11 10 12 13 9 4 3 2 1 16 15 19 20 18 17 14 29 27 30 24 22 25 28 21 23 26",
			want:  "6 11 13 12 10 9 8 7 5 4 3 2 16 20 19 18 17 15 29 30 27 24 28 25 22 26 23 21 14 1",
		},
	}

	for _, test := range tests {
		pos = 0
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%q\nwant=%q", result, test.want)
		}
	}
}
