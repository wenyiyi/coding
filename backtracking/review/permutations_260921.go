package review

/*
给定一个包含互不相同整数的数组 nums，返回它所有可能的排列

Input:
nums = [1,2,3]

Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

[]

*/

/*
todo 指针知识巩固
result 这个盒子
┌──────────────┐
│ [[1,2,3]]    │
└──────────────┘
地址：深圳市XX号

&result: 拿到盒子的地址
*result  = 地址指向的那个真正的东西

	&

东西  ─────────────→  地址

	←─────────────
	       *

外面传地址：&
里面拿东西：*
*/

/*
todo 数组和切片
最简单的辨认方法就看 [] 里面有没有数字
[3]int   → Array 数组，固定 3 个
[]int    → Slice 切片，长度可变
*/
func permute(nums []int) [][]int {
	used := make(map[int]bool, len(nums))
	path := []int{}
	result := [][]int{}

	// todo &result
	backtrack(nums, used, path, &result)
	return result
}

// todo result *[][]int
func backtrack(nums []int, used map[int]bool, path []int, result *[][]int) {

	// 一轮结束的条件
	if len(path) == len(nums) {
		// todo 不能直接保存 path，因为 slice 可能共享底层数组，因为path还需要回溯，继续变化，所以保存副本
		//              ┌───────────┐
		// path ───────→│ 1 | 2 | 3 │
		//              └───────────┘
		//                   ↑
		//                   │
		// result[0] ────────┘
		temp := append([]int{}, path...)
		// todo *result 找到地址指向的result，然后把temp放进去
		*result = append(*result, temp)
		return
	}

	// 一个一个数字去填
	for _, num := range nums { // todo for num := range nums 这里的num不是数值，而是index
		if used[num] {
			continue
		}
		// 选择
		used[num] = true
		path = append(path, num)
		// 递归
		backtrack(nums, used, path, result)
		// 撤销
		used[num] = false
		path = path[:len(path)-1]
	}

}
