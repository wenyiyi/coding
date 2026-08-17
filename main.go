package main

import "fmt"
import "coding/two_pointers"

func main() {
	nums := []int{1, 3, 4, 6, 8, 10, 13}
	target := 13
	result := two_pointers.TowSum(nums, target)
	fmt.Println(result)
}
