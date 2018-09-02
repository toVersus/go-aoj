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
			input: "8\ninsert 30\ninsert 88\ninsert 12\ninsert 1\ninsert 20\ninsert 17\ninsert 25\nprint",
			want:  " 1 12 17 20 25 30 88\n 30 12 1 20 17 25 88",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
