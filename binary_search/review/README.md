# Binary Search

## 1. Binary Search 最核心的思考方式

不要背：

- `left < right` 还是 `left <= right`
- `right = mid` 还是 `right = mid - 1`
- `left = mid` 还是 `left = mid + 1`

每次先问：

1. 我现在维护的搜索区间是什么？
2. 答案应该往左找还是往右找？
3. `mid` 还有没有可能是答案？
4. 如果 `mid` 已经确定不是答案，能不能把它排除？

---

# 2. 最重要的方向感

数组下标：

```text
← 小                              大 →

0   1   2   3   4   5   6