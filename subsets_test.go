package subsets

import(
	"testing"
	"bufio"
	"strings"
)

func TestArrayIsSubset_HashMap(t *testing.T){
	cases := []struct {
		A []string
		B []string
		R bool
		Expect bool
	}{
		{
			A: []string{"test","testing","tested"},
			B: []string{"test","tested"}, 
			R: true,
			Expect: true,
		},{
			A: []string{"test","testing","tested"},
			B: []string{"test","invalid"}, 
			R: false,
			Expect: false,
		},
	}

	for _, cItem := range cases {
		returned := ArrayIsSubset_HashMap(cItem.A, cItem.B, true, false)
		if(returned != cItem.Expect){
			t.Errorf("ArrayIsSubset_HashMap for ", cItem.A, ", ", cItem.B, " returned: ", cItem.Expect)
		}
	}
}

func TestArrayIsSubset_BruteForce(t *testing.T){
	cases := []struct {
		A []string
		B []string
		R bool
		Expect bool
	}{
		{
			A: []string{"test","testing","tested"},
			B: []string{"test","tested"}, 
			R: true,
			Expect: true,
		},{
			A: []string{"test","testing","tested"},
			B: []string{"test","invalid"}, 
			R: false,
			Expect: false,
		},
	}

	for _, cItem := range cases {
		returned := ArrayIsSubset_BruteForce(cItem.A, cItem.B, cItem.R)
		if(returned != cItem.Expect){
			t.Errorf("ArrayIsSubset_BruteForce for ", cItem.A, ", ", cItem.B, " returned: ", cItem.Expect)
		}
	}
}

func TestArrayIsSubset_RuneTree(t *testing.T){
	cases := []struct {
		A []string
		B []string
		R bool
		Expect bool
	}{
		{
			A: []string{"test","testing","tested"},
			B: []string{"test","tested"}, 
			R: true,
			Expect: true,
		},{
			A: []string{"test","testing","tested"},
			B: []string{"test","invalid"}, 
			R: false,
			Expect: false,
		},
	}

	for _, cItem := range cases {
		returned := ArrayIsSubset_RuneTree(cItem.A, cItem.B, cItem.R)
		if(returned != cItem.Expect){
			t.Errorf("ArrayIsSubset_RuneTree for ", cItem.A, ", ", cItem.B, " returned: ", cItem.Expect)
		}
	}
}

func TestReadScannerWords(t *testing.T){

	cases := []struct {
		S *bufio.Scanner
		Expect []string
	}{
		{
			S: bufio.NewScanner(strings.NewReader("Testing\t this for 1,000 testing\n purposes 0:)\r")),
			Expect: []string{"testing","this","for","testing","purposes"},
		},
	}

	for _, cItem := range cases {
		words := ReadScannerWords(cItem.S)
		returned := ArrayIsSubset_BruteForce(words, cItem.Expect, true)
		
		if(returned == false){
			t.Errorf("ReadScannerWords returned: %t", returned)
		}
	}
}