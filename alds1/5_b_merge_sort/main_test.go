package main

import (
	"reflect"
	"strings"
	"testing"
)

var mergeSortTests = []struct {
	name string
	text string
	want string
}{
	{
		name: "should return stably sorted int slice",
		text: "10\n8 5 9 2 6 3 7 1 10 4\n",
		want: "1 2 3 4 5 6 7 8 9 10\n34",
	},
	{
		name: "should return original int slice because it has already been sorted",
		text: "10\n1 2 3 4 5 6 7 8 9 10\n",
		want: "1 2 3 4 5 6 7 8 9 10\n34",
	},
	{
		name: "should return stably sorted int slice from large input",
		text: "200\n0 33 43 62 29 0 8 52 56 56 19 11 51 43 5 8 93 30 66 69 32 17 47 72 68 80 23 49 92 64 69 51 27 90 24 35 20 44 10 62 84 63 1 10 36 76 31 29 97 75 91 90 44 34 25 29 30 27 26 43 34 4 60 49 20 56 32 72 13 90 9 19 5 95 49 27 19 97 24 96 49 56 84 93 45 7 6 9 54 52 65 83 38 1 90 30 37 95 56 63 11 27 42 6 68 12 1 10 80 58 71 31 14 47 64 97 25 38 31 18 87 51 87 13 79 95 50 50 13 62 34 73 47 21 30 40 57 78 26 3 97 8 93 88 38 85 93 88 20 11 46 87 10 9 87 68 3 73 0 74 5 83 52 70 87 2 10 5 10 0 96 42 85 60 47 24 31 58 86 19 0 15 55 82 74 61 6 2 68 65 70 18 23 89 6 19 30 55 32 93\n",
		want: "0 0 0 0 0 1 1 1 2 2 3 3 4 5 5 5 5 6 6 6 6 7 8 8 8 9 9 9 10 10 10 10 10 10 11 11 11 12 13 13 13 14 15 17 18 18 19 19 19 19 19 20 20 20 21 23 23 24 24 24 25 25 26 26 27 27 27 27 29 29 29 30 30 30 30 30 31 31 31 31 32 32 32 33 34 34 34 35 36 37 38 38 38 40 42 42 43 43 43 44 44 45 46 47 47 47 47 49 49 49 49 50 50 51 51 51 52 52 52 54 55 55 56 56 56 56 56 57 58 58 60 60 61 62 62 62 63 63 64 64 65 65 66 68 68 68 68 69 69 70 70 71 72 72 73 73 74 74 75 76 78 79 80 80 82 83 83 84 84 85 85 86 87 87 87 87 87 88 88 89 90 90 90 90 91 92 93 93 93 93 93 95 95 95 96 96 97 97 97 97\n1544",
	},
}

func TestMergeSort(t *testing.T) {
	for _, testcase := range mergeSortTests {
		t.Log(testcase.name)

		cnt = 0
		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}
	}
}

var result string

func BenchmarkMergeSort(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		for _, testcase := range mergeSortTests {
			b.StopTimer()
			cnt = 0
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			str = answer(r)
		}
	}
	result = str
}
