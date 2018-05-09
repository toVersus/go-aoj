package main

import (
	"os"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		file string
		text string
		want string
	}{
		{
			name: "should print the sorted hand",
			file: "in1.txt",
			text: `5
H4 C9 S4 D2 C3`,
			want: "D2 C3 H4 S4 C9",
		},
		{
			name: "should print the hand without sorting",
			file: "in2.txt",
			text: `2
S1 H1`,
			want: "S1 H1",
		},
		{
			name: "should print the simple sorted hand",
			file: "in3.txt",
			text: `2
S2 D1`,
			want: "D1 S2",
		},
		{
			name: "should print the sorted hand, from large input",
			file: "in10.txt",
			text: `36
H8 S5 C5 H9 S4 D9 S7 S2 C7 S8 H6 C9 C3 C1 S9 S1 H3 D2 D7 D3 D8 H7 S6 H4 C2 H2 D5 H1 D4 C4 S3 D1 C6 D6 C8 H5`,
			want: "C1 S1 H1 D1 S2 D2 C2 H2 C3 H3 D3 S3 S4 H4 D4 C4 S5 C5 D5 H5 H6 S6 C6 D6 S7 C7 D7 H7 H8 S8 D8 C8 H9 D9 C9 S9",
		},
	}

	for _, testcase := range tests {
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
		cards, N := newCards(f)
		f.Close()

		result := bubbleSort(cards, N).String()
		if !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name      string
		file      string
		text      string
		want      string
		stability string
	}{
		{
			name: "should print the sorted hand, but unstable",
			file: "in1.txt",
			text: `5
H4 C9 S4 D2 C3`,
			want:      "D2 C3 S4 H4 C9",
			stability: "Not stable",
		},
		{
			name: "should print the hand without sorting",
			file: "in2.txt",
			text: `2
S1 H1`,
			want:      "S1 H1",
			stability: "Stable",
		},
		{
			name: "should print the simple sorted hand, but stable",
			file: "in3.txt",
			text: `2
S2 D1`,
			want:      "D1 S2",
			stability: "Stable",
		},
		{
			name: "should print the sorted hand, but unstable, from large input",
			file: "in10.txt",
			text: `36
H8 S5 C5 H9 S4 D9 S7 S2 C7 S8 H6 C9 C3 C1 S9 S1 H3 D2 D7 D3 D8 H7 S6 H4 C2 H2 D5 H1 D4 C4 S3 D1 C6 D6 C8 H5`,
			want:      "C1 S1 H1 D1 S2 D2 C2 H2 C3 H3 D3 S3 H4 S4 D4 C4 D5 C5 S5 H5 S6 C6 D6 H6 S7 D7 H7 C7 H8 S8 D8 C8 S9 C9 H9 D9",
			stability: "Not stable",
		},
	}

	for _, testcase := range tests {
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
		cards, N := newCards(f)
		f.Close()

		selection := selectionSort(cards, N)
		result := selection.String()
		if !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v,\n want => %#v", result, testcase.want)
		}

		if stability := isStable(selection, bubbleSort(cards, N)); !reflect.DeepEqual(stability, testcase.stability) {
			t.Errorf("result => %#v,\n want => %#v", stability, testcase.stability)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
