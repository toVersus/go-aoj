package main

import (
	"os"
	"reflect"
	"testing"
)

var linkingTests = []struct {
	name string
	file string
	text string
	want string
}{
	{
		name: "should return single key",
		file: "in1.txt",
		text: "9\ninsert 5\ninsert 2\ninsert 3\ninsert 1\ndelete 3\ninsert 6\ndelete 5\ndeleteFirst\ndeleteLast\n",
		want: "1",
	},
	{
		name: "should insert same key too many times",
		file: "in3.txt",
		text: "39\ninsert 1\ninsert 2\ninsert 1\ninsert 1\ninsert 1\ninsert 2\ninsert 2\ninsert 1\ndelete 2\ndeleteLast\ndeleteFirst\ninsert 2\ninsert 2\ninsert 2\ninsert 2\ninsert 2\ninsert 2\ninsert 2\ninsert 2\ninsert 2\ninsert 2\ndeleteFirst\ndeleteFirst\ndelete 2\ndelete 2\ndeleteFirst\ndelete 1\ninsert 3\ninsert 1\ninsert 1\ndelete 2\ndelete 2\ndelete 2\ndeleteLast\ninsert 4\ndelete 2\ndeleteLast\ndeleteFirst\ndelete 2\n",
		want: "1 1 3 2 1",
	},
	{
		name: "should delete non-existing key",
		file: "in4.txt",
		text: "8\ninsert 1000000000\ninsert 999999999\ndeleteLast\ninsert 1234566890\ninsert 5\ndeleteFirst\ninsert 7\ndelete 5\n",
		want: "7 1234566890 999999999",
	},
	{
		name: "should return keys from large input",
		file: "in8.txt",
		text: "111\ninsert 8\ninsert 7\ninsert 7\ninsert 8\ninsert 7\ninsert 8\ninsert 1\ninsert 1\ninsert 1\ninsert 1\ninsert 1\ninsert 1\ninsert 1\ninsert 1\ninsert 8\ninsert 8\ninsert 1\ninsert 1\ninsert 8\ndeleteLast\ninsert 5\ninsert 1\ninsert 1\ninsert 7\ninsert 8\ndeleteFirst\ninsert 8\ninsert 7\ninsert 8\ninsert 1\ninsert 1\ninsert 1\ninsert 1\ndelete 8\ninsert 1\ninsert 1\ninsert 7\ninsert 1\ninsert 1\ninsert 1\ndelete 1\ninsert 1\ninsert 8\ninsert 5\ndelete 1\ndelete 1\ninsert 7\ndelete 8\ndelete 7\ninsert 8\ninsert 1\ninsert 7\ninsert 5\ninsert 7\ninsert 1\ninsert 1\ninsert 1\ninsert 1\ninsert 8\ndeleteFirst\ninsert 8\ninsert 1\ninsert 1\ninsert 5\ninsert 7\ninsert 1\ninsert 1\ninsert 8\ninsert 1\ndeleteFirst\ninsert 7\ninsert 8\ndelete 1\ninsert 7\ndelete 7\ninsert 1\ninsert 1\ninsert 8\ninsert 7\ninsert 8\ninsert 5\ninsert 7\ndeleteLast\ndeleteLast\ninsert 8\ndeleteFirst\ninsert 8\ndelete 1\ninsert 7\ndelete 8\ninsert 8\ninsert 5\ndeleteLast\ninsert 7\ninsert 1\ninsert 1\ninsert 8\ninsert 13\ndeleteLast\ndeleteLast\ninsert 8\ndelete 8\ninsert 13\ndelete 1\ndelete 1\ndeleteLast\ndelete 1\ninsert 1\ndelete 1\ninsert 1\ndelete 1\n",
		want: "13 13 8 7 5 8 7 7 5 8 7 8 8 7 8 1 7 5 1 1 8 1 1 1 1 7 5 7 1 8 5 1 7 1 1 1 1 1 1 7 8 7 1 1 5 8 1 1 8 8 1 1 1 1 1 1 1",
	},
}

func TestGetLists(t *testing.T) {
	for _, testcase := range linkingTests {
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

		if err := runInst(f); err != nil {
			t.Error(err)
		}
		f.Close()

		if result := stringList(); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkGetLists(b *testing.B) {
	for _, testcase := range linkingTests {
		f, err := os.Create(testcase.file)
		if err != nil {
			b.Errorf("could not create a file: %s\n  %s", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range linkingTests {
			b.StopTimer()
			f, _ := os.Open(testcase.file)
			b.StartTimer()

			if err := runInst(f); err != nil {
				b.Error(err)
			}
			f.Close()
			stringList()
		}
	}
	b.StopTimer()

	for _, testcase := range linkingTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
