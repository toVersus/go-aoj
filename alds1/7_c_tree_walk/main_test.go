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
			input: "1\n0 -1 -1", want: "Preorder\n 0\nInorder\n 0\nPostorder\n 0",
		},
		{
			input: `5
3 4 0
4 -1 -1
1 -1 -1
2 -1 -1
0 1 2`,
			want: `Preorder
 3 4 0 1 2
Inorder
 4 3 1 0 2
Postorder
 4 1 2 0 3`,
		},
		{
			input: `4
1 0 -1
0 2 -1
2 3 -1
3 -1 -1`,
			want: `Preorder
 1 0 2 3
Inorder
 3 2 0 1
Postorder
 3 2 0 1`,
		},
		{
			input: `9
0 1 4
1 2 3
2 -1 -1
3 -1 -1
4 5 8
5 6 7
6 -1 -1
7 -1 -1
8 -1 -1`,
			want: `Preorder
 0 1 2 3 4 5 6 7 8
Inorder
 2 1 3 0 6 5 7 4 8
Postorder
 2 3 1 6 7 5 8 4 0`,
		},
		{
			input: `24
0 -1 -1
1 2 5
2 3 4
3 -1 -1
4 -1 -1
5 6 7
6 8 11
7 -1 21
8 9 -1
9 -1 10
10 -1 -1
11 -1 -1
12 20 -1
13 -1 1
14 -1 13
15 -1 14
16 -1 -1
17 16 15
18 17 -1
19 18 -1
20 19 -1
21 -1 22
22 23 0
23 -1 -1`,
			want: `Preorder
 12 20 19 18 17 16 15 14 13 1 2 3 4 5 6 8 9 10 11 7 21 22 23 0
Inorder
 16 17 15 14 13 3 2 4 1 9 10 8 6 11 5 7 21 23 22 0 18 19 20 12
Postorder
 16 3 4 2 10 9 8 11 6 23 0 22 21 7 5 1 13 14 15 17 18 19 20 12`,
		},
	}

	for _, test := range tests {
		out.Reset()
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%q\nwant=%q", result, test.want)
		}
	}
}
