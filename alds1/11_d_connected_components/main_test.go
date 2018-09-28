package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestConnectedComponent(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "2 0\n1\n0 1\n",
			want:  "no",
		},
		{
			input: "10 9\n0 1\n0 2\n3 4\n5 7\n5 6\n6 7\n6 8\n7 8\n8 9\n3\n0 1\n5 9\n1 3\n",
			want:  "yes\nyes\nno",
		},
		{
			input: "10 5\n0 1\n1 2\n2 9\n3 4\n6 7\n12\n0 1\n0 2\n0 9\n3 4\n6 7\n7 6\n0 5\n0 6\n3 7\n3 5\n5 8\n8 5\n",
			want:  "yes\nyes\nyes\nyes\nyes\nyes\nno\nno\nno\nno\nno\nno",
		},
		{
			input: "30 27\n0 1\n2 3\n3 4\n4 5\n5 6\n7 8\n8 9\n9 10\n10 11\n11 12\n12 13\n13 14\n14 15\n15 16\n16 17\n17 18\n18 19\n19 20\n20 21\n21 22\n22 23\n23 24\n24 25\n25 26\n26 27\n27 28\n28 29\n20\n12 25\n1 12\n13 21\n7 22\n13 14\n3 11\n15 26\n9 26\n26 29\n18 23\n24 27\n7 16\n15 18\n4 14\n9 20\n3 23\n7 14\n2 17\n6 9\n18 22\n",
			want:  "yes\nno\nyes\nyes\nyes\nno\nyes\nyes\nyes\nyes\nyes\nyes\nyes\nno\nyes\nno\nyes\nno\nno\nyes",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
