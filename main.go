package main

import (
	"coding/sliding_window"
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

	nums := []int{2, 1, 5, 1, 3, 2}
	sum := sliding_window.MaxSum(nums, 3)
	fmt.Println(sum)
}
