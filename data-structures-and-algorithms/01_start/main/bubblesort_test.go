package main

import "testing"

var array = []int{5, 2, 7, 4, 6, 3, 1}
var result = []int{1, 2, 3, 4, 5, 6, 7}

func TestBubbleSort(t *testing.T) {
	BubbleSort(array, len(array))
	for i, val := range array {
		if val != result[i] {
			t.Fail()
		}
	}
}
