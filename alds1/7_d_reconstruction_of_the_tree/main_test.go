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
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%q\nwant=%q", result, test.want)
		}
	}
}
