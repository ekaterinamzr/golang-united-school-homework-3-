package homework

import "testing"

type testPairSlice struct {
	arr      []int64
	reversed []int64
}

func TestReverse(t *testing.T) {
	var test testPairSlice = testPairSlice{[]int64{1, 2, 3, 4, 5}, []int64{5, 4, 3, 2, 1}}

	res := reverse(test.arr)

	for i, _ := range res {
		if res[i] != test.reversed[i] {
			t.Error("Expected ", test.reversed,
				"got ", res)
		}
	}

}

type testPairMap struct {
	input  map[int]string
	sorted []string
}

func TestSortMapValues(t *testing.T) {
	var test testPairMap = testPairMap{map[int]string{2: "a", 0: "b", 1: "c"},
		[]string{"b", "c", "a"}}

	res := sortMapValues(test.input)

	for i, _ := range res {
		if res[i] != test.sorted[i] {
			t.Error("Expected ", test.sorted,
				"got ", res)
		}
	}

}
