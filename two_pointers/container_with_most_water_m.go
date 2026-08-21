package two_pointers

/*
	DESCRIPTION (inspired by Leetcode.com)
	Given an array heights where each element represents the height of a vertical line,
    pick two lines to act as the walls of a container.
    Return the maximum area (amount of water) the container can hold.

    What is area? Width × height, where width is the distance between walls,
    and height is the shorter wall (water overflows at the shorter wall).

    input: heights = [3, 4, 1, 2, 2, 4, 1, 3, 2]
    output: 21  # walls at indices 0 and 7 (both height 3): width=7, height=3, area=21
*/

/*
	What is area? Just width × height:
	Width: How far apart the two walls are (right_index - left_index)
	Height: The shorter wall's height (min(heights[left], heights[right]))
*/

func MaxArea(heights []int) int {
	left, right := 0, len(heights)-1
	maxArea := 0

	for left < right {
		height := min(heights[left], heights[right])
		with := right - left
		currArea := with * height
		if currArea > maxArea {
			maxArea = currArea
		}
		// move the shorter wall's pointer
		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
