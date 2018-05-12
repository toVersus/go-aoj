package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

var arithmeticTests = []struct {
	name string
	text string
	want int
}{
	{
		name: "should add the two operands",
		text: "1 1 +",
		want: 2,
	},
	{
		name: "should multiply two operands",
		text: "1 2 + 3 4 + *",
		want: 21,
	},
	{
		name: "should get the sum of numbers from 1 to 10",
		text: "1 2 3 4 5 6 7 8 9 10 + + + + + + + + +",
		want: 55,
	},
	{
		name: "should get the sum of numbers from 1 to 10",
		text: "1000 500 132 + 132 - 2 * 100 10 10 * + - 50 50 50 50 + + + + 10 999 * + 24 + + 24 - 10 * 15000 -",
		want: 104900,
	},
}

func TestCompute(t *testing.T) {
	for _, testcase := range arithmeticTests {
		t.Log(testcase.name)
		r := strings.NewReader(testcase.text)
		op := []string{}
		sc := bufio.NewScanner(r)
		if sc.Scan() {
			op = strings.Fields(sc.Text())
		}

		if result := compute(op); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}
	}
}

func BenchmarkCompute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range arithmeticTests {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			op := []string{}
			sc := bufio.NewScanner(r)
			if sc.Scan() {
				op = strings.Fields(sc.Text())
			}
			compute(op)
		}
	}
}
