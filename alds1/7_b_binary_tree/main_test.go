package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "1\n0 -1 -1", want: "node 0: parent = -1, sibling = -1, degree = 0, depth = 0, height = 0, root"},
		{
			input: `5
3 4 0
4 -1 -1
1 -1 -1
2 -1 -1
0 1 2`,
			want: `node 0: parent = 3, sibling = 4, degree = 2, depth = 1, height = 1, internal node
node 1: parent = 0, sibling = 2, degree = 0, depth = 2, height = 0, leaf
node 2: parent = 0, sibling = 1, degree = 0, depth = 2, height = 0, leaf
node 3: parent = -1, sibling = -1, degree = 2, depth = 0, height = 2, root
node 4: parent = 3, sibling = 0, degree = 0, depth = 1, height = 0, leaf`},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%q, want=%q", result, test.want)
		}
	}
}
