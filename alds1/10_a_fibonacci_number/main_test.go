package main

import (
	"reflect"
	"testing"
)

func TestFibonacciNumber(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 1,
			want:  "1",
		},
		{
			input: 15,
			want:  "987",
		},
		{
			input: 29,
			want:  "832040",
		},
	}

	for _, test := range tests {
		if result := answer(test.input); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
