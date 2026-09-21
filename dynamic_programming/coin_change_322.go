package dynamic_programming

/*
https://leetcode.com/problems/coin-change/

Coin Change
Return the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite(无限的) number of each kind of coin.


Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0

DP 规律
1. Can I break the problem into smaller subproblems?
2. Does the same subproblem appear multiple times?
3. Same state → same answer?
*/

/*
todo
问题：凑出 amount，最少需要多少枚硬币

DP四问
第一步：怎么定义dp[]
DP 第一件事就是把最终问题泛化：
凑出 11 最少需要几枚？
        ↓
凑出任意金额 i 最少需要几枚？
dp[i] = 凑出金额 i 最少需要的硬币数量
dp[0] = 凑 0 元最少几枚
dp[1] = 凑 1 元最少几枚
dp[2] = 凑 2 元最少几枚
...
dp[11] = 凑 11 元最少几枚

第二步：要得到dp[i]，最后一步怎么来
最后拿 1 元：
dp[i-1] + 1

最后拿 2 元：
dp[i-2] + 1

最后拿 5 元：
dp[i-5] + 1

dp[i] = min(dp[i], dp[i-coin]+1)

第三步：初始化状态
dp[0]=0 ✅
❌
dp[1] = 0
dp[2] = 0
dp[3] = 0
一个常见办法就是初始化成一个足够大的、不可能成为正确答案的数
如：假设有 1 元硬币，凑 amount 元 最多也只需 amount 枚硬币
dp[i] = amount + 1

第四步：return 什么
return dp[amount]

(该题特殊点：如果 dp[amount] == amount + 1
说明它从初始的“不可能值”开始，经过所有状态转移后依然没有被更新，即 amount 无法凑出。
return -1)
*/

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}

	// 0 1 2 3 4 5 6 7 8 9 10 11   0～amount一共需要amount+1个位置
	maxCount := amount + 1
	// []int{} 创建长度为0的slice，不能dp[0]=0
	// 第一步定义：dp[i] = 凑出金额 i 最少需要多少枚硬币
	dp := make([]int, maxCount)

	// 第三步初始化： dp[0] = 0，其他位置设置成一个很大的数字，不可能的数字
	for i := 1; i < maxCount; i++ {
		dp[i] = maxCount
	}

	// 一开始只有 1元硬币
	// coin = 当前硬币 (1,2,5)
	// currAmount = 当前金额
	// todo coin在外层，先拿coin，再去凑
	for _, coin := range coins {
		// 然后分别用1元凑 1，2 ，3，4，5
		for currAmount := coin; currAmount <= amount; currAmount++ {
			// coins = [1, 2, 5]
			// currAmount = 11
			// 第二步 要得到dp[i]，最后一步怎么来
			dp[currAmount] = min(dp[currAmount], dp[currAmount-coin]+1)
		}
	}

	// dp[amount]初始=amount+1,如果一直没有被重新赋值表示没有任何一种硬币组合能组成总金额，返回-1
	if dp[amount] == maxCount {
		return -1
	}

	// 第四步，返回
	return dp[amount]
}
