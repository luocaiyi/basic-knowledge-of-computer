package main

// SumI 数组求和：迭代
func SumI(array []int) int {
	sum := 0
	for _, val := range array {
		sum += val
	}
	return sum
}

// SumDecrease 减而治之
func SumDecrease(array []int, n int) int {
	if n < 1 {
		return 0
	} else {
		return SumDecrease(array, n-1) + array[n-1]
	}
}

// SumDivide 分而治之
func SumDivide(array []int, lo int, hi int) int {
	if lo == hi {
		return array[lo]
	}
	mi := (lo + hi) >> 1
	return SumDivide(array, lo, mi) + SumDivide(array, mi+1, hi)
}
