package dp

/*
https://www.scribd.com/document/1059776896/Expedia-SDE2-OA-Coding%E5%A4%A7%E5%85%A8?utm_source=chatgpt.com
Grouping Options(分组方案)

Given a number of people n and a number of groups k,(给定人数和组数)
determine(求出/确定) the distinct ways(不同的方案) to form(形成) k contiguous(邻近的，相连的) groups from the n people
while adhering(遵守) to the following conditions:
(给定 n 个人和 k 个组，求把 n 个人分成 k 组一共有多少种不同的分组方案)


The sum of all group sizes(所有组的人数之和) must equal the total number of people.
Each group size must be greater than or equal to the size of the group to its left.(每一组的人数必须 大于或等于 它左边那一组的人数，即递增)
The group formations(分组方案) must be distinct, meaning they must differ in at least one group.
For example, [1, 1, 1, 3] is distinct from [1, 1, 1, 2] but not from [1, 3, 1, 1].

Example
people = 8
groups = 4
The 5 distinct options to form 4 groups with 8 people under the rules are:
[1, 1, 1, 5] = 8
[1, 1, 2, 4] = 8
[1, 1, 3, 3] = 8
[1, 2, 2, 3] = 8
[2, 2, 2, 2] = 8
In each option, the groups are distinct and each group's size is greater than or equal to the group to its left.

Function Description
Complete the function countOptions in the editor with the following parameters:
int people: an integer that denotes(表示) the number of people in the row
int groups: an integer that denotes the number of groups to form

Returns
long int: the number of ways that n participants(参与者) can be divided into k groups that satisfy the conditions
*/

/*
num1 + num2 + num3 + num4 = people
num1 <= num2 <= num3 <= num4

把 people 拆成 groups 个正整数，并且这几个整数必须非递减，求有多少种拆法
*/

/*
把 DP 表想成一个答案查询表
              分几组 g
         0   1    2    3
0        ┌────────────────
1个人    │    1    0    0
2个人    │    1    1    0
3个人    │    1    1    1
4个人    │    1    2    1
5个人    │    1    2    2
6个人    │    1    3    3


	  dp[5][2]
      5个人分2组
          ↓
    ┌─────┴─────┐
    ↓           ↓
  [1,4]       [2,3]
    ↓           ↓
  去掉1       每组减1
    ↓           ↓
   [4]         [1,2]
    ↓           ↓
dp[4][1]     dp[3][2]

dp[5][2] = dp[4][1] + dp[3][2] = dp[5-1][2-1] + dp[5-2][2] = dp[p][g] = dp[p-g][g] + dp[p-1][g-1]

*/

func countOptions(people int32, groups int32) int64 {
	// 把 people 个人拆成 groups 组有多少种方案
	//dp
	//↓
	//[
	//  nil,   // dp[0] 0个人
	//  nil,   // dp[1] 1个人
	//  nil,   // dp[2]
	//  ...
	//  nil    // dp[n] 8个人   长度：8+1=9
	//]
	dp := make([][]int64, people+1) // people+1 表示行数
	for i := range dp {
		// 注意：每一行还要继续初始化，第二维表示的是组数 g，而我们需要表示
		// 	0组
		//	1组
		//	2组
		//	...
		//	groups组
		dp[i] = make([]int64, groups+1)
	}

	// 初始化
	dp[0][0] = 1

	for p := 1; p <= int(people); p++ {
		for g := 1; g <= int(groups); g++ {
			// 公式：dp[p][g] = dp[p-g][g] + dp[p-1][g-1]
			// 防止 p-g < 0，导致数组下标越界
			if (p - g) >= 0 {
				dp[p][g] += dp[p-g][g]
			}
			dp[p][g] += dp[p-1][g-1]
		}
	}
	return dp[people][groups]
}
