package dynamic_programming

/*
Coin Change II

https://leetcode.com/problems/coin-change-ii/
You are given an integer array coins representing coins of different denominations(面额)
and an integer amount representing a total amount of money.
Return the number of combinations that make up that amount.

If that amount of money cannot be made up by any combination(组合) of the coins, return 0.

You may assume that you have an infinite number of each kind of coin.(每种面额的硬币都有无限个，可以重复使用)
The final answer is guaranteed to fit into a signed 32-bit integer.

Example 1:
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

Example 2:
Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.

Example 3:
Input: amount = 10, coins = [10]
Output: 1
*/

/*
问题：凑出 amount，一共有多少种硬币组合

DP四问
第一步：怎么定义dp[]
DP 第一件事就是把最终问题泛化：
dp[i]= 凑出金额 i, 一共有多少种硬币组合


第二步：要得到 dp[i]，怎么从更小的问题得到？
coins = [1,2,5]
选择 1 元硬币：
dp[i-1] 种方案

选择 2 元硬币：
dp[i-2] 种方案

选择 5 元硬币：
dp[i-5] 种方案

dp[i] += dp[i-coin]


第三步：初始化状态
dp[0]=1 凑出金额 0 有1种方法，什么都不拿
其他位置可以初始化为 0，表示暂时还没找到任何组合


第四步：return 什么
return dp[amount]
凑不出来时 dp[amount] 自然保持 0

*/

/*
总结

LC322 Coin Change
dp[i] = 最少硬币数
转移：min
→ 不统计方案数量
→ 顺序不会造成重复计数
→ coin 外层 / amount 外层都可以


LC518 Coin Change II
dp[i] = 组合数量
转移：+
→ 会统计每一种方案
→ 顺序会影响“组合还是排列”
→ 求组合必须 coin 外层

组合：不看顺序，1+2 和 2+1 是同一个
排列：看顺序，1+2 和 2+1 是两个

先固定 coin，其实就是人为规定一个使用硬币的顺序，避免一会儿生成 [1,2]，一会儿又生成 [2,1]
*/

func change(amount int, coins []int) int {
	if len(coins) == 0 {
		return 0
	}

	maxLen := amount + 1
	// amount = 5, coins = [1,2,5]
	// 第一步：定义 dp[i] = 凑出金额 i 有多少种组合     0 1 2 ... 5
	dp := make([]int, maxLen)

	// 第三步：初始化，凑出金额 0 有1种方法，什么都不拿
	dp[0] = 1

	// 一开始只有 1元硬币
	// coin = 当前硬币
	// currAmount = 当前金额
	// todo coin在外层，先拿coin，再去凑amount
	for _, coin := range coins {
		// 然后分别用1元凑 1，2 ，3，4，5
		for currAmount := coin; currAmount <= amount; currAmount++ {
			// 第二步：要得到 dp[i]，怎么从更小的问题得到
			dp[currAmount] += dp[currAmount-coin]
		}

	}

	// 第四步返回
	return dp[amount]
}
