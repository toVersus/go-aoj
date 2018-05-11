package main

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"
)

var arithmeticTests = []struct {
	name string
	file string
	text string
	want int
}{
	{
		name: "should add the two operands",
		file: "in1.txt",
		text: "1 1 +",
		want: 2,
	},
	{
		name: "should multiply two operands",
		file: "in2.txt",
		text: "1 2 + 3 4 + *",
		want: 21,
	},
	{
		name: "should get the sum of numbers from 1 to 10",
		file: "in6.txt",
		text: "1 2 3 4 5 6 7 8 9 10 + + + + + + + + +",
		want: 55,
	},
	{
		name: "should get the sum of numbers from 1 to 10",
		file: "in9.txt",
		text: "1000 500 132 + 132 - 2 * 100 10 10 * + - 50 50 50 50 + + + + 10 999 * + 24 + + 24 - 10 * 15000 -",
		want: 104900,
	},
}

func TestCompute(t *testing.T) {
	for _, testcase := range arithmeticTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		f, err = os.Open(testcase.file)
		if err != nil {
			t.Errorf("could not open a file: %s\n  %s", testcase.file, err)
		}

		op := []string{}
		sc := bufio.NewScanner(f)
		if sc.Scan() {
			op = strings.Fields(sc.Text())
		}
		f.Close()

		if result := compute(op); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkCompute(b *testing.B) {
	for _, testcase := range arithmeticTests {
		f, _ := os.Create(testcase.file)
		f.WriteString(testcase.text)
		f.Close()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range arithmeticTests {
			b.StopTimer()
			f, _ := os.Open(testcase.file)
			b.StartTimer()

			op := []string{}
			sc := bufio.NewScanner(f)
			if sc.Scan() {
				op = strings.Fields(sc.Text())
			}
			f.Close()
			compute(op)
		}
	}
	b.StopTimer()

	for _, testcase := range arithmeticTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
