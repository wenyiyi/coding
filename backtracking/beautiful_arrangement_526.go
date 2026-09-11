package backtracking

/*
https://leetcode.com/problems/beautiful-arrangement/description/
Beautiful Arrangement

Suppose you have n integers labeled 1 through n.
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
*/
/*
perm[i] % i == 0
i % perm[i] == 0

n=2
[1,2]
i=1
perm[1]=1
1%1=0

i=2
perm[2]=2
2%2=0

[2,1]
i=1
perm[1]=2
2%1=0

i=2
perm[2]=2
2%2=0

n=3
[1, 2, 3]
i=1: 1 % 1 == 0 ✅
i=2: 2 % 2 == 0 ✅
i=3: 3 % 3 == 0 ✅

[1,3,2]
i=1: 1 % 1 == 0 ✅
i=2:
3 % 2 != 0
2 % 3 != 0


This is basically a permutation problem with pruning.
I use backtracking to fill each position,
and I only continue if the current number satisfies the divisibility condition.

一格一格放数字
position 1
   ↓
选一个没用过的数字
   ↓
position 2
   ↓
再选一个没用过的数字
   ↓
position 3
每次选择之前先判断,num % position == 0 || position % num == 0
*/

/*
为什么i不需要重新回到1？
因为 i 不是你手动修改后需要恢复的状态，它是每一层递归自己的局部变量，

i            → 每层自己的变量 → return 后自然恢复
used[]       → 所有层共享状态 → 必须手动回溯


i=1
│
├─ 选择 num=1
│    │
│    └─ i=2
│         │
│         └─ i=3
│              return
│         ↑
│       回到 i=2
│    ↑
│  回到 i=1
│
├─ 选择 num=2
│    │
│    └─ i=2
│
*/

func countArrangement(n int) int {
	// n+1，因为下标从1开始
	// n=3
	// 	[1,2,3]
	// 0 1 2 3   长度4
	// 同一个 i 要共用一个 []
	perm := make([]bool, n+1)
	return gerValidPerm(1, n, perm)
}

func gerValidPerm(i int, n int, perm []bool) int {
	// 所有位置都放完了 [ , , ]，表示找到1个合法的
	if i > n {
		return 1
	}

	var result = 0
	// [ , , ] 然后一格一格放数字，确定一个之后，就 判断下一个 i+1
	for num := 1; num <= n; num++ {

		if perm[num] {
			continue
		}

		// 不符合条件
		if num%i != 0 && i%num != 0 {
			continue
		}

		// 做选择
		perm[num] = true
		// 递归放下一个数字, 共用一个 perm。 只是把 i + 1 的值传给下一层，上一层的 i 根本没变
		result += gerValidPerm(i+1, num, perm)
		// 撤销选择
		perm[num] = false
	}

	return result
}
