package main

import (
	"reflect"
	"strings"
	"testing"
)

var linearSearchTests = []struct {
	name string
	text string
	want int
}{
	{
		name: "should return correct number of union of slices",
		text: "5\n1 2 3 4 5\n3\n3 4 1\n",
		want: 3,
	},
	{
		name: "no intersection between two slices",
		text: "10\n10 9 8 7 6 5 4 3 2 1\n5\n11 12 13 14 15\n",
		want: 0,
	},
	{
		name: "should return correct number of union of slices from large input",
		text: "200\n0 933 743 262 529 700 508 752 256 256 119 711 351 843 705 108 393 330 366 169 932 917 847 972 868 980 223 549 592 164 169 551 427 190 624 635 920 944 310 862 484 363 301 710 236 876 431 929 397 675 491 190 344 134 425 629 30 727 126 743 334 104 760 749 620 256 932 572 613 490 509 119 405 695 49 327 719 497 824 596 649 356 184 93 245 7 306 509 754 352 665 783 738 801 690 330 337 195 656 963 11 427 42 106 968 212 1 510 480 658 571 331 814 847 564 197 625 438 931 18 487 151 187 913 179 995 750 750 913 562 134 273 547 521 830 140 557 678 726 503 597 408 893 988 238 85 93 188 720 211 746 387 710 209 887 668 103 473 900 674 105 183 952 370 787 302 410 905 110 400 996 142 585 860 47 924 731 158 386 219 400 415 55 682 874 61 6 602 268 365 470 518 723 89 106 319 130 655 732 993\n30\n0 25 28 30 60 71 73 78 79 83 94 95 116 132 143 155 167 199 221 237 248 264 267 299 328 331 334 338 369 376\n",
		want: 4,
	},
}

func TestLinearSearch(t *testing.T) {
	for _, testcase := range linearSearchTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	b.ResetTimer()
	for _, testcase := range linearSearchTests {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			answer(r)
		}
	}
}
