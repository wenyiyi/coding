# Dynamic Programming

## DP 四问

### 1. dp 表示什么？

把题目的「最终问题」泛化成「任意中间状态」。

例如 LC62：

最终问题：

    到 (m-1, n-1) 有多少条路径？

泛化：

    到任意 (i, j) 有多少条路径？

所以：

    dp[i][j] = 从 (0,0) 到 (i,j) 的路径数量


### 2. 当前状态从哪里来？

问：

> 当前答案可以由哪些更小的答案组成？

常见思考方式：

- Grid：最后一步从哪里来？
- Sequence：以 i 结尾，前一个状态是谁？
- Knapsack：选择 / 不选择当前物品？
- Counting：所有方案可以拆成哪些互斥情况？
- Bitmask：加入一个还没有使用的元素。


### 3. 初始状态是什么？

找到最小问题 / 空状态。

例如：

    Unique Paths:
    dp[i][0] = 1
    dp[0][j] = 1

    Grouping Options:
    dp[0][0] = 1


### 4. 最终答案在哪里？

不一定总是 dp[n]。

可能是：

    dp[n]
    dp[n-1]
    dp[m-1][n-1]
    max(dp)
    dp[fullMask]


---

# Pattern Map

## 1. Grid / Position DP

### LC62 Unique Paths

定义：

    dp[i][j] = 从 (0,0) 到 (i,j) 的路径数量

转移：

    dp[i][j]
    = dp[i-1][j] + dp[i][j-1]

思考：

    当前格子的最后一步从哪里来？
    → 上面 + 左边


---

## 2. Sequence DP

### LC53 Maximum Subarray

定义：

    dp[i] = 以 nums[i] 结尾的最大连续子数组和

转移：

    dp[i] = max(
        nums[i],
        dp[i-1] + nums[i]
    )

思考：

    到 nums[i]：
    → 自己重新开始
    → 接在前面的答案后面


### LC300 Longest Increasing Subsequence

定义：

    dp[i] = 以 nums[i] 结尾的最长递增子序列长度

思考：

    枚举 j < i

    如果 nums[j] < nums[i]
    → nums[i] 可以接在 nums[j] 后面


---

## 3. Knapsack DP

### LC322 Coin Change

定义：

    dp[x] = 凑出金额 x 所需要的最少硬币数

关键词：

    最少
    → min


### LC518 Coin Change II

定义：

    dp[x] = 凑出金额 x 的组合数量

关键词：

    多少种方案
    → +


LC322 vs LC518：

    LC322：求最优值 → min
    LC518：求方案数 → +


---

## 4. Counting / Classification DP

### Grouping Options

定义：

    dp[p][g]
    = p 个人分成 g 个非递减正数组的方案数量

把所有答案分成两类：

    最小组 = 1
        → 删除这个 1
        → dp[p-1][g-1]

    最小组 > 1
        → 每组减 1
        → dp[p-g][g]

所以：

    dp[p][g]
    = dp[p-1][g-1] + dp[p-g][g]


---

## 5. Bitmask DP

### LC1986 Minimum Number of Work Sessions

核心：

    用一个整数的二进制位表示
    “哪些任务已经完成”。

例如：

    mask = 0101

表示：

    task 0 ✓
    task 1 ✗
    task 2 ✓
    task 3 ✗

识别信号：

    n 比较小
    +
    需要记录哪些元素已经选择 / 完成

    → 考虑 Bitmask DP


---

# DP Pattern Recognition

看到 DP 题，先不要想公式。

先问：

    1. 最终问题是什么？
                ↓
    2. 泛化成任意中间状态
                ↓
    3. dp[state] 表示什么？
                ↓
    4. 当前状态从哪些更小状态来？
                ↓
    5. Base Case
                ↓
    6. Final Answer


常见入口：

Grid:
到当前位置的答案是什么？

Sequence:
以 i 结尾的答案是什么？

Knapsack:
凑出容量 / 金额 x 的答案是什么？

Counting:
所有方案能不能拆成几个互斥情况？

Bitmask:
哪些元素已经使用 / 完成？