package sorting_by_end_time

import "sort"

/*
https://www.hellointerview.com/learn/code/intervals/non-overlapping-intervals

Non-Overlapping Intervals
DESCRIPTION (inspired by Leetcode.com)
Write a function to return the minimum number of intervals that must be removed from a given array intervals,
where intervals[i] consists of a starting point starti and an ending point endi,
to ensure that the remaining intervals do not overlap.
Intervals that only touch at their endpoints are not considered overlapping (e.g., [2,5] and [5,7] do not overlap).

Input: intervals = [[1,3],[5,8],[4,10],[11,13]]
Output: 1
Explanation: Removing the interval [4,10] leaves all other intervals non-overlapping.

*/

func nonOverlappingIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 1   3
	//           5     8
	//       4             10
	//                         11   13
	// sort by end time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// record current end
	end := intervals[0][1]
	count := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			end = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count
}
