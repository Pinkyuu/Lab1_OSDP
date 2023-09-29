package main

import (
	"testing"
)

func stringSplitBlocksTest(t *testing.T) {

	testTable := []struct {
		a        string
		expected []string
	}{{
		a:        "Akaka, 325, 25.08.2002",
		expected: []string{"Akaka", "325", "25.08.2002"},
	},
		{
			a:        "asdafga, hdhfhdfh, hdfhgdfg, dhdfhfdh",
			expected: []string{"asdafga", "hdhfhdfh", "hdfhgdfg", "dhdfhfdh"},
		},
		{
			a:        "",
			expected: []string{},
		},
	}

	for i, testCase := range testTable {
		result := StringSplitBlocks(testCase.a)
		if result[i] != testCase.expected[i] {
			//fmt.Printf("Incorrect result. Expect %d, got %d", testCase.expected[i], result[i])
			t.Log("Incorrect result")
		} else {
			t.Log("Correct result")
		}
	}
}

/*func addBirdTest()

func addFishTest()

func addInsectsTest()

func addtest()

func remtest()

func printest()*/
