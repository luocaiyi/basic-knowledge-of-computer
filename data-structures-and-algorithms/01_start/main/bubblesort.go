package main

func BubbleSort(array []int, n int) {
	for sorted := false; !sorted; n-- {
		sorted = true
		for i := 1; i < n; i++ {
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
				sorted = false
			}
		}
	}
}
