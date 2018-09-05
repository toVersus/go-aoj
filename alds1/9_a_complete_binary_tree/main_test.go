package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCompleteBinaryTree(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "5\n7 8 1 2 3",
			want:  "node 1: key = 7, left key = 8, right key = 1, \nnode 2: key = 8, parent key = 7, left key = 2, right key = 3, \nnode 3: key = 1, parent key = 7, \nnode 4: key = 2, parent key = 8, \nnode 5: key = 3, parent key = 8, ",
		},
		{
			input: "17\n1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17",
			want:  "node 1: key = 1, left key = 2, right key = 3, \nnode 2: key = 2, parent key = 1, left key = 4, right key = 5, \nnode 3: key = 3, parent key = 1, left key = 6, right key = 7, \nnode 4: key = 4, parent key = 2, left key = 8, right key = 9, \nnode 5: key = 5, parent key = 2, left key = 10, right key = 11, \nnode 6: key = 6, parent key = 3, left key = 12, right key = 13, \nnode 7: key = 7, parent key = 3, left key = 14, right key = 15, \nnode 8: key = 8, parent key = 4, left key = 16, right key = 17, \nnode 9: key = 9, parent key = 4, \nnode 10: key = 10, parent key = 5, \nnode 11: key = 11, parent key = 5, \nnode 12: key = 12, parent key = 6, \nnode 13: key = 13, parent key = 6, \nnode 14: key = 14, parent key = 7, \nnode 15: key = 15, parent key = 7, \nnode 16: key = 16, parent key = 8, \nnode 17: key = 17, parent key = 8, ",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
