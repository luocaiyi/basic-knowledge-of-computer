package main

func Hailstone(n int) []int {
	hailstone := []int{n}
	for n > 1 {
		if n%2 == 1 {
			n = 3*n + 1
		} else {
			n = n / 2
		}
		hailstone = append(hailstone, n)
	}
	return hailstone
}
