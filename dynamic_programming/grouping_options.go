package dynamic_programming

/*
https://www.scribd.com/document/1059776896/Expedia-SDE2-OA-Coding%E5%A4%A7%E5%85%A8?utm_source=chatgpt.com
(traveloka笔试题)

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
问题：people 个人分成 groups 组，每组人数 > 0，并且人数非递减，一共有多少种分法？

第一步：怎么定义dp[][]
dp[i][j] = 把 i 个人分成 j 个非递减的正整数分组，有多少种分法

第二步：要得到 dp[i][j]，怎么从更小的问题得到？
拿具体值去推出规律 dp[5][2]

	     分几组 g
	0        1    2    3

0        ┌────────────────
1个人    │    1    0    0
2个人    │    1    1    0
3个人    │    1    1    1
4个人    │    1    2    1
5个人    │    1    2    2
6个人    │    1    3    3

dp[5][2] = 5个人2组 = (1,4)(2,3) = dp[4][1]+dp[3][2]
dp[i][j] = dp[i-1][j-1]+dp[i-j][j] todo dp[i-j][j] 而不是 dp[i-2]
把所有合法答案按照“第一组是不是1”分成两类：

 1. 第一组 = 1
    删除第一组
    → dp[i-1][j-1]

 2. 第一组 > 1
    因为非递减，所以所有组都 > 1
    每组减1，总人数减少j，组数不变
    → dp[i-j][j]

第三步：如何初始化
dp[0][0]=1 0个人分成0组有1种方案，什么都不分
dp[0][1]=0
dp[1][0]=0
所以其他都可以初始化为0

第四部：return 什么
return dp[people][groups]
*/
func countOptions(people int32, groups int32) int64 {

	//dp
	//↓
	//[
	//  nil,   // dp[0] 0个人
	//  nil,   // dp[1] 1个人
	//  nil,   // dp[2]
	//  ...
	//  nil    // dp[n] 8个人   长度：8+1=9
	//]
	// 把 p 个人拆成 g 组有多少种方案
	dp := make([][]int64, people+1) // 长度 people+1，下标才能到 people
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
