package backtracking

/*

https://leetcode.com/problems/permutations/description/
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]

Constraints:
1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

/*
path [] 一格一格填，直到 len(path) = len(nums)
加入 result
然后撤销，重新填
做一个选择

	↓

递归

	↓

撤销刚刚这个选择
*/
func permute(nums []int) [][]int {
	// nums 中的数字互不相同，所以可以直接用 value 记录是否使用
	used := make(map[int]bool)
	path := []int{}
	result := [][]int{}

	backtrack1(nums, used, path, &result)

	return result
}

func backtrack1(
	nums []int,
	used map[int]bool,
	path []int,
	result *[][]int,
) {
	// 终止条件：已经选了所有数字
	if len(path) == len(nums) {
		// todo path 后面还会被回溯修改，所以保存一份 copy
		temp := append([]int{}, path...) // todo 把一个 slice 的所有元素作为多个参数传给 append 时，需要 ...
		*result = append(*result, temp)
		return
	}

	// 当前这一层尝试所有数字
	for _, num := range nums {
		// 当前 path 已经使用过
		if used[num] {
			continue
		}

		// 1. 做选择
		used[num] = true
		path = append(path, num)

		// 2. 递归
		backtrack1(nums, used, path, result)

		// 3. 撤销选择
		path = path[:len(path)-1]
		used[num] = false
	}
}

/*
LC46 Permutations vs LC526 Beautiful Arrangement

LC46:
- 状态：path + used
- 当前递归层：继续往 path 放一个数字
- 不需要 position：
  len(path) 已经表示当前递归深度
- 需要 path：最终要保存具体的排列
- used：防止同一个数字重复使用
- 基本没有额外 pruning
- 终止条件：
  len(path) == len(nums)

LC526:
- 状态：position + used
- 当前递归层：给第 position 个位置选择一个数字
- 需要 position：合法条件依赖当前位置 position
- 不需要 path：只统计合法排列数量，不需要保存具体排列
- used：防止同一个数字重复使用
- pruning：
  num % position == 0 || position % num == 0
- 终止条件：
  position > n

核心区别：
LC46 需要“具体排列” → path
LC526 需要“当前位置”做合法性判断 → position

共同模板：
for 每个候选 {
    if 已使用 {
        continue
    }

    if 不满足条件 { // 部分题才有 pruning
        continue
    }

    做选择
    backtrack(...)
    撤销选择
}
*/
