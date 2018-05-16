package main

import (
	"reflect"
	"strings"
	"testing"
)

var allocateCapacityTests = []struct {
	name string
	text string
	want string
}{
	{
		name: "should print the minimum value of max capacity",
		text: "5 3\n8\n1\n7\n3\n9\n",
		want: "10",
	},
	{
		name: "should print the minimum value of max capacity from middle input",
		text: "64 8\n6\n7\n4\n15\n13\n3\n2\n2\n4\n14\n5\n3\n4\n6\n4\n5\n1\n3\n2\n7\n2\n7\n5\n8\n3\n5\n3\n7\n7\n4\n7\n4\n3\n4\n7\n9\n4\n4\n1\n6\n5\n2\n6\n3\n1\n7\n11\n4\n7\n4\n6\n5\n5\n7\n7\n7\n7\n4\n7\n7\n6\n3\n4\n7\n",
		want: "45",
	},
}

func TestAllocateCapacity(t *testing.T) {
	for _, testcase := range allocateCapacityTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
	}
}
