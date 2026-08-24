package main

import (
	"coding/sliding_window/fixed_length_sliding_window"
	"fmt"
)

func main() {
	//nums := []int{1, 3, 4, 6, 8, 10, 13}
	//target := 13
	//towSum := two_pointers.TowSum(nums, target)
	//fmt.Println(towSum)

	//heights := []int{3, 4, 1, 2, 2, 4, 1, 3, 2}
	//maxArea := two_pointers.MaxArea(heights)
	//fmt.Println(maxArea)

	//nums := []int{11, 4, 9, 6, 15, 18}
	//result := two_pointers.TriangleNumber(nums)
	//fmt.Println(result)

	nums := []int{2, 11, 4, 5, 3, 9, 2}
	sum := fixed_length_sliding_window.MaxScore(nums, 3)
	fmt.Println(sum)
}
