package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestMatrixChainMultiplication(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "6\n30 35\n35 15\n15 5\n5 10\n10 20\n20 25\n",
			want:  "15125",
		},
		{
			input: "1\n1 34",
			want:  "0",
		},
		{
			input: "50\n1 34\n34 44\n44 13\n13 30\n30 1\n1 9\n9 3\n3 7\n7 7\n7 20\n20 12\n12 2\n2 44\n44 6\n6 9\n9 44\n44 31\n31 17\n17 20\n20 33\n33 18\n18 48\n48 23\n23 19\n19 31\n31 24\n24 50\n50 43\n43 15\n15 20\n20 2\n2 28\n28 41\n41 25\n25 36\n36 21\n21 45\n45 11\n11 13\n13 35\n35 14\n14 2\n2 11\n11 37\n37 27\n27 32\n32 30\n30 48\n48 26\n26 42\n",
			want:  "28807",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
