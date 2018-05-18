package main

import (
	"reflect"
	"strings"
	"testing"
)

var searchTests = []struct {
	name string
	text string
	want string
}{
	{
		name: "should return the possibility to make integer by adding combination of element",
		text: "5\n1 5 7 10 21\n10\n2 4 6 15 17 8 22 21 100 35",
		want: "no\nno\nyes\nyes\nyes\nyes\nyes\nyes\nno\nno",
	},
	{
		name: "should return the possibility to make integer by adding combination of element from large input",
		text: "20\n1 34 44 63 30 1 9 53 57 57 20 12 52 44 6 9 94 31 67 70\n200\n1933 918 1848 1973 869 981 224 550 1593 165 170 552 428 191 625 636 921 1945 1311 1863 485 1364 1302 711 1237 1877 432 1930 398 676 492 191 1345 135 426 1630 31 728 1127 1744 1335 105 761 1750 1621 257 933 1573 1614 491 510 1120 1406 1696 50 328 720 1498 825 597 650 1357 185 1094 246 1008 1307 510 755 353 1666 784 739 802 1691 1331 1338 196 657 1964 12 1428 43 1107 969 213 1002 511 481 659 572 1332 815 848 565 1198 1626 439 932 19 488 152 1188 1914 180 996 1751 1751 914 1563 1135 274 1548 522 831 1141 1558 1679 727 504 598 409 894 1989 1239 1086 94 1189 721 1212 747 388 1711 1210 1888 669 1104 1474 1901 1675 1106 184 953 371 1788 1303 1411 1906 111 1401 1997 143 586 1861 1048 1925 1732 1159 1387 1220 401 1416 56 1683 1875 1062 1007 1603 1269 1366 471 519 724 90 1107 320 1131 1656 733 1994 1975 1596 811 1674 1056 731 96 246 706 1695 949 1874 354 475 761 742 757 1645 1144 1641\n",
		want: "no\nno\nno\nno\nno\nno\nyes\nyes\nno\nyes\nyes\nyes\nyes\nyes\nyes\nyes\nno\nno\nno\nno\nyes\nno\nno\nyes\nno\nno\nyes\nno\nyes\nyes\nyes\nyes\nno\nyes\nyes\nno\nyes\nyes\nno\nno\nno\nyes\nno\nno\nno\nyes\nno\nno\nno\nyes\nyes\nno\nno\nno\nyes\nyes\nyes\nno\nno\nyes\nyes\nno\nyes\nno\nyes\nno\nno\nyes\nno\nyes\nno\nno\nyes\nno\nno\nno\nno\nyes\nyes\nno\nyes\nno\nyes\nno\nno\nyes\nno\nyes\nyes\nyes\nyes\nno\nno\nno\nyes\nno\nno\nyes\nno\nyes\nyes\nyes\nno\nno\nyes\nno\nno\nno\nno\nno\nno\nyes\nno\nyes\nno\nno\nno\nno\nyes\nyes\nyes\nyes\nno\nno\nno\nno\nyes\nno\nyes\nno\nyes\nyes\nno\nno\nno\nyes\nno\nno\nno\nno\nno\nyes\nno\nyes\nno\nno\nno\nno\nyes\nno\nno\nyes\nyes\nno\nno\nno\nno\nno\nno\nno\nyes\nno\nyes\nno\nno\nno\nno\nno\nno\nno\nyes\nyes\nyes\nyes\nno\nyes\nno\nno\nyes\nno\nno\nno\nno\nno\nno\nyes\nyes\nyes\nyes\nno\nno\nno\nyes\nyes\nno\nyes\nno\nno\nno\nno",
	},
}

func TestExhaustiveSearch(t *testing.T) {
	for _, testcase := range searchTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
	}
}
