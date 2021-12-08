package main

import "testing"

var arrayMap = map[int][]int{
	55: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	13: {5, 8},
}

func TestSumI(t *testing.T) {
	for result, array := range arrayMap {
		if SumI(array) != result {
			t.Fatalf("SumIErr sum(%v) != %d", array, result)
		}
	}
}

func TestSumDecrease(t *testing.T) {
	for result, array := range arrayMap {
		if SumDecrease(array, len(array)) != result {
			t.Fatalf("SumDecrease sum(%v) != %d", array, result)
		}
	}
}

func TestSumDivide(t *testing.T) {
	for result, array := range arrayMap {
		if SumDivide(array, 0, len(array)-1) != result {
			t.Fatalf("SumDivide sum(%v) != %d", array, result)
		}
	}
}
