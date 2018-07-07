package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestRootTrees(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "1\n0 0", want: "node 0: parent = -1, depth = 0, root, []"},
		{
			input: `4
0 0
1 3 2 3 0
2 0
3 0`,
			want: `node 0: parent = 1, depth = 1, leaf, []
node 1: parent = -1, depth = 0, root, [2, 3, 0]
node 2: parent = 1, depth = 1, leaf, []
node 3: parent = 1, depth = 1, leaf, []`},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%q, want=%q", result, test.want)
		}
	}
}
