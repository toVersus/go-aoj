package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestBinarySearchTree2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "10\ninsert 30\ninsert 88\ninsert 12\ninsert 1\ninsert 20\nfind 12\ninsert 17\ninsert 25\nfind 16\nprint\n",
			want:  "yes\nno\n 1 12 17 20 25 30 88\n 30 12 1 20 17 25 88",
		},
		{
			input: "24\ninsert 8\ninsert 2\ninsert 3\ninsert 7\ninsert 22\ninsert 1\ninsert 10\ninsert 9\ninsert 15\ninsert 13\nfind 8\nfind 2\nfind 3\nfind 7\nfind 22\nfind 1\nfind 10\nfind 9\nfind 15\nfind 13\nfind 4\nfind 5\nfind100\nprint\n",
			want:  "yes\nyes\nyes\nyes\nyes\nyes\nyes\nyes\nyes\nyes\nno\nno\n 1 2 3 7 8 9 10 13 15 22\n 8 2 1 3 7 22 10 9 15 13",
		},
		{
			input: "28\ninsert 20\ninsert 19\ninsert 50\nfind 61\nfind -100\nfind 20\nprint\ninsert 1000000000\ninsert 1\ninsert 18\ninsert 5\ninsert 3\ninsert 7\ninsert 30\ninsert 25\ninsert 40\ninsert 41\ninsert 42\nfind 3\nfind 7\nfind 42\nfind 20\nfind 1000000000\nfind 23\nfind 2\nfind 60\nfind 100\nprint\n",
			want:  "no\nno\nyes\n 19 20 50\n 20 19 50\nyes\nyes\nyes\nyes\nyes\nno\nno\nno\nno\n 1 3 5 7 18 19 20 25 30 40 41 42 50 1000000000\n 20 19 1 18 5 3 7 50 30 25 40 41 42 1000000000",
		},
	}

	for _, test := range tests {
		out.Reset()
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
