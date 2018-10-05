package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestAllPairsShortestPath(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "4 6\n0 1 1\n0 2 5\n1 2 2\n1 3 4\n2 3 1\n3 2 7\n",
			want:  "0 1 3 4\nINF 0 2 3\nINF INF 0 1\nINF INF 7 0",
		},
		{
			input: "4 6\n0 1 1\n0 2 -5\n1 2 2\n1 3 4\n2 3 1\n3 2 7\n",
			want:  "0 1 -5 -4\nINF 0 2 3\nINF INF 0 1\nINF INF 7 0",
		},
		{
			input: "4 6\n0 1 1\n0 2 5\n1 2 2\n1 3 4\n2 3 1\n3 2 -7\n",
			want:  "NEGATIVE CYCLE",
		},
		{
			input: "20 30\n0 1 7\n1 2 37\n2 3 4\n3 4 0\n4 5 41\n5 6 47\n6 7 48\n7 8 3\n8 9 49\n9 10 37\n10 11 15\n11 12 29\n12 13 49\n13 14 10\n14 15 20\n15 16 31\n16 17 2\n17 18 42\n18 19 23\n19 0 28\n1 17 31\n10 4 13\n11 9 7\n6 14 46\n10 12 6\n15 6 18\n5 18 21\n9 18 32\n2 8 41\n0 13 6\n",
			want:  "0 7 44 48 48 89 54 102 85 134 171 186 177 6 16 36 67 38 80 103\n124 0 37 41 41 82 129 177 78 127 164 179 170 130 140 160 191 31 73 96\n117 124 0 4 4 45 92 140 41 90 127 142 133 123 133 153 184 155 66 89\n113 120 157 0 0 41 88 136 139 188 225 240 231 119 129 149 180 151 62 85\n113 120 157 161 0 41 88 136 139 188 225 240 231 119 129 149 180 151 62 85\n72 79 116 120 120 0 47 95 98 147 184 199 190 78 88 108 139 110 21 44\n183 190 227 231 150 191 0 48 51 100 137 152 143 189 46 66 97 99 132 155\n135 142 179 183 102 143 189 0 3 52 89 104 95 141 151 171 202 173 84 107\n132 139 176 180 99 140 186 234 0 49 86 101 92 138 148 168 199 170 81 104\n83 90 127 131 50 91 137 185 168 0 37 52 43 89 99 119 150 121 32 55\n105 112 149 153 13 54 101 149 152 22 0 15 6 55 65 85 116 118 54 77\n90 97 134 138 57 98 126 174 175 7 44 0 29 78 88 108 139 128 39 62\n205 212 249 253 247 288 97 145 148 197 234 249 0 49 59 79 110 112 154 177\n156 163 200 204 198 239 48 96 99 148 185 200 191 0 10 30 61 63 105 128\n146 153 190 194 188 229 38 86 89 138 175 190 181 152 0 20 51 53 95 118\n126 133 170 174 168 209 18 66 69 118 155 170 161 132 64 0 31 33 75 98\n95 102 139 143 143 184 149 197 180 229 266 281 272 101 111 131 0 2 44 67\n93 100 137 141 141 182 147 195 178 227 264 279 270 99 109 129 160 0 42 65\n51 58 95 99 99 140 105 153 136 185 222 237 228 57 67 87 118 89 0 23\n28 35 72 76 76 117 82 130 113 162 199 214 205 34 44 64 95 66 108 0",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		if result := answer(r); !reflect.DeepEqual(result, test.want) {
			t.Errorf("\ngot =%q\nwant=%q", result, test.want)
		}
	}
}
