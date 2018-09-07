package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "insert 8\ninsert 2\nextract\ninsert 10\nextract\ninsert 11\nextract\nextract\nend\n",
			want:  "8\n10\n11\n2",
		},
		{
			input: "insert 120\ninsert 88\ninsert 3\ninsert 999999999\ninsert 7\nextract\ninsert 100\ninsert 21\ninsert 100\ninsert 100\ninsert 10\nextract\ninsert 100\nextract\ninsert 777\ninsert 100\nextract\nextract\nextract\nextract\nextract\ninsert 120\nextract\nextract\nend\n",
			want:  "999999999\n120\n100\n777\n100\n100\n100\n100\n120\n88",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
