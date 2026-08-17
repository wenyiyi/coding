package main

import (
	"coding/two_pointers"
	"fmt"
)

func main() {
	//nums := []int{1, 3, 4, 6, 8, 10, 13}
	//target := 13
	//towSum := two_pointers.TowSum(nums, target)
	//fmt.Println(towSum)

	heights := []int{3, 4, 1, 2, 2, 4, 1, 3, 2}
	maxArea := two_pointers.MaxArea(heights)
	fmt.Println(maxArea)
}
