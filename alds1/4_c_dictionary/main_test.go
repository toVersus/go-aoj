package main

import (
	"reflect"
	"strings"
	"testing"
)

var dictionaryTests = []struct {
	name string
	text string
	want string
}{
	{
		name: "simple insert and find instruction sets",
		text: "15\ninsert AAA\ninsert AAC\ninsert AGA\ninsert AGG\ninsert TTT\nfind AAA\nfind CCC\nfind CCC\ninsert CCC\nfind CCC\nfind CC\ninsert T\nfind TTT\nfind T\nfind A\n",
		want: "yes\nno\nno\nyes\nno\nyes\nyes\nno",
	},
	{
		name: "random insert and find instruction sets",
		text: "20\nfind CTT\nfind TAG\ninsert AA\ninsert CTC\nfind G\ninsert TGT\nfind AAC\nfind TTG\ninsert AA\ninsert TTG\ninsert G\ninsert A\ninsert GG\ninsert C\ninsert TT\nfind T\ninsert C\nfind GG\nfind G\ninsert CG\n",
		want: "no\nno\nno\nno\nno\nno\nyes\nyes",
	},
	{
		name: "should get the correct yes/no response from large input",
		text: "100\ninsert TGCAAA\nfind TTTCACG\ninsert ACTAAAT\nfind CTTGATAAGG\ninsert CGAATCCTTG\nfind CGTGTGAACA\nfind CGCTCTC\ninsert CAACAACC\nfind GACTGCGG\ninsert ATTTGGAA\ninsert GTTGTAC\nfind GTTTCT\nfind CGGCTC\nfind GGTCAC\nfind CAATGTGCT\ninsert CAGCTAGTG\ninsert GAAGCA\nfind GGTATTGG\nfind ACGGTCG\nfind ACGTGCTGTC\nfind CCGACA\ninsert ACTTCCGTCT\ninsert GCAATCCT\ninsert CACATGAG\ninsert GAGCGAA\ninsert TGTGCCA\nfind CTTGTAG\nfind TAACAGGCAA\ninsert TCGTCCTGCA\ninsert TCCGAGCAAC\ninsert TCTAGTCTAG\ninsert TTCGGCG\nfind CCCGGG\ninsert ATACGGCCC\ninsert GCATCT\ninsert AGTCTAACAC\nfind ATCAACG\ninsert TCTGTGATTG\nfind AACGGTGT\nfind GTACAAACG\nfind AGTAAGGG\ninsert TTTATAGAA\nfind TTTAAA\ninsert TACCCCCGC\nfind GCCACGAC\nfind AGGCTT\nfind GGGATA\nfind TCCATT\ninsert AGGGGA\ninsert TGTCACTTAA\nfind GCAGAAAC\nfind AGTCAAGCT\nfind GACCCAGATC\ninsert TTTACAGGC\nfind GATGAC\ninsert GGCGCCCG\nfind CTATTA\nfind TTGTGGGGG\nfind ATAATAACCG\nfind TCATGCAACC\ninsert TGTTTATA\nfind GCCCCCATC\ninsert AGTGAG\nfind GGCACTC\ninsert AAGTCAG\nfind GCTTAG\ninsert TTTCCCAT\ninsert CCCTCC\nfind GTATAG\ninsert GCCTGACGA\nfind CTGGTGAT\nfind TAACAGC\nfind ACATCTCT\ninsert GTTATGTCC\ninsert GCAGGAGTG\ninsert GATGTT\ninsert TGCATGTTG\ninsert CGACTG\ninsert CCCCCGTCT\ninsert TCCACAT\nfind TGGACG\nfind TATAGATGC\nfind CTCCAT\ninsert TTTTCAACAA\ninsert GGCGGTAACT\ninsert TGCGATGA\nfind ATTGAATCC\nfind GGTCGAT\nfind GGATCCCAT\ninsert CGTGCA\nfind CTCAAAGGA\ninsert GGCTAGC\nfind TGTACGAT\nfind CTTTCTTAA\ninsert AGATTCCCGC\ninsert AACGCTGC\ninsert TGTCTGAT\nfind CGCACT\nfind ATCGGTG\nfind CTGGAGTA\n",
		want: "no\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno\nno",
	},
}

func TestDictionary(t *testing.T) {
	for _, testcase := range dictionaryTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := answer(r); !reflect.DeepEqual(result, testcase.want) {
			t.Errorf("result => %#v\n, want => %#v", result, testcase.want)
		}
	}
}

func BenchmarkDictionary(b *testing.B) {
	b.ResetTimer()
	for _, testcase := range dictionaryTests {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			answer(r)
		}
	}
}
