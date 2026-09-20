package backtracking

/*
https://leetcode.com/problems/beautiful-arrangement/description/
(traveloka笔试题)
Beautiful Arrangement

Suppose you have n integers labeled(已标记) 1 through n.
A permutation(排列) of those n integers perm (1-indexed) is considered a beautiful arrangement(安排)
if for every i (1 <= i <= n), either of the following is true(以下任一说法正确，及or):

perm[i] is divisible by i.
i is divisible by perm[i].

Given an integer n, return the number of the beautiful arrangements that you can construct(构造).

consecutive 连续的
at least 至少
either 任何一个


Example 1:
Input: n = 2
Output: 2
Explanation:
The first beautiful arrangement is [1,2]:
    - perm[1] = 1 is divisible by i = 1
    - perm[2] = 2 is divisible by i = 2
The second beautiful arrangement is [2,1]:
    - perm[1] = 2 is divisible by i = 1
    - i = 2 is divisible by perm[2] = 1

Example 2:
Input: n = 1
Output: 1

Backtracking 关键词：constructed
*/

/*

Pattern: Backtracking + Pruning剪枝

本质：
Permutation + 额外合法条件

每层递归填一个 position：
1. 遍历所有数字
2. 已使用 → skip
3. 不满足整除条件 → pruning
4. 做选择 used[num] = true
5. 递归填下一个 position
6. 撤销 used[num] = false

position = 1 → 给第 1 个位置选数字
position = 2 → 给第 2 个位置选数字
position = 3 → 给第 3 个位置选数字

合法条件：
num % position == 0 || position % num == 0

todo 只求合法排列数量，不需要具体排列，因此不需要 path。

注意：
position → 每层递归自己的局部状态，return 后自然恢复
used[]   → 所有递归层共享，必须手动撤销
*/

func countArrangement(n int) int {
	// n+1，因为下标从1开始
	// n=3
	// 	[1,2,3]
	// 0 1 2 3   长度4
	// 同一个 i 要共用一个 []
	// used := make([]bool, n+1)，用 []bool 会比 map 更轻量，因为数字范围明确就是 1...n

	// 统一用map
	used := make(map[int]bool)
	result := 0
	backtrack(1, n, used, &result)
	return result
}

// todo 注释写“为什么”，少写“代码正在做什么”
func backtrack(position int, n int, used map[int]bool, result *int) {
	// 终止条件，所有位置都放完了 [ , , ]，表示找到1个合法的
	if position > n { // i (1 <= i <= n)
		*result++
		return
	}

	for num := 1; num <= n; num++ {
		// 当前数字已使用
		if used[num] {
			continue
		}

		// 不符合当前位置的整除条件->pruning
		if num%position != 0 && position%num != 0 {
			continue
		}

		// 做选择
		used[num] = true
		// 填下一个位置
		backtrack(position+1, n, used, result)
		// 撤销刚才的选择
		// todo position 是每层自己的局部状态，不需要像 used 一样手动恢复
		// todo used[]   → 所有递归层共享，必须手动撤销
		used[num] = false
	}

	return
}
