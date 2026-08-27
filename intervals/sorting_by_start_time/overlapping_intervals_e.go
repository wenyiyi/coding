package sorting_by_start_time

import "sort"

/*
https://www.hellointerview.com/learn/code/intervals/can-attend-meetings

Can Attend Meetings
DESCRIPTION (inspired by Leetcode.com)
Write a function to check if a person can attend all the meetings scheduled without any time conflicts.
Given an array intervals, where each element [s1, e1] represents a meeting starting at time s1 and ending at time e1,
determine if there are any overlapping meetings.
If there is no overlap between any meetings, return true; otherwise, return false.

Note that meetings ending and starting at the same time, such as (0,5) and (5,10), do not conflict.

Input: intervals = [(1,5),(3,9),(6,8)]
Output: false
Explanation: The meetings (1,5) and (3,9) overlap.

Input: intervals = [(10,12),(6,9),(13,15)]
Output: true
Explanation: There are no overlapping meetings, so the person can attend all.
*/
func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}
	// sort by each first number
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		// compare the start time of the current interval with the end time of the previous interval
		// 1     5
		//   3       9
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
