package dp

/*
https://leetcode.com/problems/coin-change/

Coin Change
Return the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.


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

*/

/*
问最少需要多少枚硬币
dp[i] = 凑出金额 i 最少需要多少枚硬币
dp[0] = 0 凑出 0 元，需要 0 枚硬币
但是其他位置不能初始化成 0, dp[5] = 0, 就变成了“凑 5 元只需要 0 个硬币”，显然不对。
所以一开始把其他位置设置成一个很大的数字，不可能的数字
这里最简单可以直接,amount + 1, 因为最坏情况下，即使全用 1 元硬币，最多也就 amount 枚
*/

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}

	maxCount := amount + 1

	// coins = [1,2,5], amount = 11
	// dp[i] = 凑出金额 i 最少需要多少枚硬币
	dp := make([]int, maxCount)

	// 0 1 2 3 4 5 6 7 8 9 10 11

	//初始化： dp[0] = 0，其他位置设置成一个很大的数字，不可能的数字
	for i := 1; i < maxCount; i++ {
		dp[i] = maxCount
	}

	// 一开始只有 1元硬币
	// coin = 当前硬币
	// currAmount = 当前金额
	for _, coin := range coins {
		// 然后分别用1元凑 1，2 ，3，4，5
		for currAmount := coin; currAmount <= amount; currAmount++ {
			// coins = [1, 2, 5]
			// currAmount = 11
			// coin = 5
			// 假设我们决定拿一个 5 元硬币，问题就变成： 凑出 6 元最少需要几个硬币？
			dp[currAmount] = min(dp[currAmount], dp[currAmount-coin]+1)
		}
	}

	// dp[amount]初始=amount+1,如果一直没有被重新赋值表示没有任何一种硬币组合能组成总金额，返回-1
	if dp[amount] == maxCount {
		return -1
	}

	return dp[amount]
}
