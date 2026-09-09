package dp

/*
Coin Change II

https://leetcode.com/problems/coin-change-ii/
You are given an integer array coins representing coins of different denominations
and an integer amount representing a total amount of money.
Return the number of combinations that make up that amount.
If that amount of money cannot be made up by any combination of the coins, return 0.
You may assume that you have an infinite number of each kind of coin.
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
dp[i] = 凑出金额 i 有多少种方法
金额      0  1  2  3  4  5
dp       [?, ?, ?, ?, ?, ?]

dp[0]=1 因为凑出金额0有一种方法就是什么硬币都不拿

现在只有 1 元硬币
amount = 1
1 = 0 + 1
dp[1] += dp[0] =1

amount = 2
2 = 1 + 1
dp[2] += dp[1] = 1

金额      0  1  2  3  4  5
dp       [1, 1, 1, 1, 1, 1]

现在允许 1元 2元

我要凑 5，现在拿一个 2，那剩下是不是只需要凑 3？所以凑 3 有多少种方法，就能给凑 5 新增多少种方法
dp[a] += dp[a-coin]
*/
func change(amount int, coins []int) int {

	// amount = 5, coins = [1,2,5]
	// dp[i] = 凑出金额 i 有多少种方法     0 1 2 ... 5
	dp := make([]int, amount+1)

	// 初始化，凑出金额 0 有1种方法，什么都不拿
	dp[0] = 1

	// 一开始只有 1元
	for _, coin := range coins {

		// 然后分别用1元凑 1，2 ，3，4，5
		for currAmount := coin; currAmount <= amount; currAmount++ {
			dp[currAmount] += dp[currAmount-coin]
		}

	}

	return dp[amount]
}
