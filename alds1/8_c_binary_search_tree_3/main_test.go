package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestBinarySearchTree3(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "18\ninsert 8\ninsert 2\ninsert 3\ninsert 7\ninsert 22\ninsert 1\nfind 1\nfind 2\nfind 3\nfind 4\nfind 5\nfind 6\nfind 7\nfind 8\nprint\ndelete 3\ndelete 7\nprint\n",
			want:  "yes\nyes\nyes\nno\nno\nno\nyes\nyes\n 1 2 3 7 8 22\n 8 2 1 3 7 22\n 1 2 8 22\n 8 2 1 22",
		},
		{
			input: "43\ninsert 30\ninsert 17\ninsert 88\ninsert 53\ninsert 5\ninsert 20\ninsert 18\ninsert 28\ninsert 27\ninsert 60\nprint\nfind -1\nfind 2\nfind 3\nfind 4\nfind 5\nfind 6\nfind 10\nfind 17\nfind 28\nfind 29\nfind 30\nfind 31\nfind 50\nfind 51\nfind 52\nfind 53\nfind 54\nfind 60\nfind 88\nfind 89\ninsert 2000000000\ninsert 55\ninsert 63\ninsert -1\ninsert 8\nprint\ndelete 53\ndelete 2000000000\ndelete 20\ndelete 5\ndelete 8\nprint\n",
			want:  " 5 17 18 20 27 28 30 53 60 88\n 30 17 5 20 18 28 27 88 53 60\nno\nno\nno\nno\nyes\nno\nno\nyes\nyes\nno\nyes\nno\nno\nno\nno\nyes\nno\nyes\nyes\nno\n -1 5 8 17 18 20 27 28 30 53 55 60 63 88 2000000000\n 30 17 5 -1 8 20 18 28 27 88 53 60 55 63 2000000000\n -1 17 18 27 28 30 55 60 63 88\n 30 17 -1 27 18 28 88 60 55 63",
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
