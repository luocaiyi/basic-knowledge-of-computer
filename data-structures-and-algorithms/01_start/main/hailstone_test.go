package main

import "testing"

var ns = []int{7, 42}

var resultMap = map[int][]int{
	7:  {7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1},
	42: {42, 21, 64, 32, 16, 8, 4, 2, 1},
}

func TestHailstone(t *testing.T) {
	for _, n := range ns {
		t.Logf("Test Hailstone(%d)", n)
		for i, val := range resultMap[n] {
			if Hailstone(n)[i] != val {
				t.Fatal("Hailstone Err")
			}
		}
		t.Log("ok")
	}
}
