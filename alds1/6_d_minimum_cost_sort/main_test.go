package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestMinCostSort(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"5\n1 5 3 4 2", 7},
		{"8\n2 1 3 4 8 5 6 7", 39},
		{"1\n1", 0},
		{"100\n1 49 37 4 23 6 7 10 9 33 20 12 13 14 5 30 61 15 19 11 21 22 18 24 53 54 27 28 29 16 97 93 89 34 72 36 3 8 39 40 64 59 98 75 45 68 47 48 35 66 17 50 25 31 90 56 57 58 42 60 79 94 63 41 65 38 67 43 69 70 71 96 73 74 32 76 77 78 51 80 81 82 83 84 85 86 87 88 55 99 91 92 2 95 62 44 26 46 52 100", 2677},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("result => %#v,\n want => %#v", result, test.want)
		}
	}
}
