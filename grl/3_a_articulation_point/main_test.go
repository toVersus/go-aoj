package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestArticulationPoint(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "1 0\n",
			want:  "",
		},
		{
			input: "2 1\n0 1\n",
			want:  "",
		},
		{
			input: "3 2\n0 1\n0 2\n",
			want:  "0\n",
		},
		{
			input: "3 2\n0 1\n1 2\n",
			want:  "1\n",
		},
		{
			input: "4 4\n0 1\n0 2\n1 2\n2 3\n",
			want:  "2\n",
		},
		{
			input: "5 4\n0 1\n1 2\n2 3\n3 4\n",
			want:  "1\n2\n3\n",
		},
		{
			input: "15 18\n0 1\n0 2\n1 2\n1 3\n3 5\n3 4\n4 6\n5 6\n5 9\n7 9\n7 8\n9 10\n10 11\n11 12\n11 13\n12 13\n13 14\n12 14\n",
			want:  "1\n3\n5\n7\n9\n10\n11\n",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
