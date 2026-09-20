package review

import "math/rand"

/*
实现 RandomizedSet 类：

Insert(val int) bool

向集合中插入 val。

如果 val 不存在，插入并返回 true
如果已经存在，返回 false
Remove(val int) bool

从集合中删除 val。

如果 val 存在，删除并返回 true
如果不存在，返回 false
GetRandom() int

随机返回集合中的一个元素。

每个元素被选中的概率相同
可以假设调用时集合中至少有一个元素

*/

type RandomizedSet struct {
	// map insert/remove
	// array getRandom
	indexMap map[int]int
	values   []int
}

func Constructor919() RandomizedSet {
	// todo 搞懂各种初始化
	// m1 := map[int]int{}       → 空 map，已初始化 ✅
	// m2 := make(map[int]int)   → 空 map，已初始化 ✅
	// var m3 map[int]int        → nil map，未初始化 ⚠️

	// []int{}            → []           len=0 → append
	// make([]int, 0)     → []           len=0 → append
	// make([]int, 3)     → [0, 0, 0]    len=3 → 可以 arr[i] =
	// var a []int        → nil slice    len=0 → 也可以 append
	return RandomizedSet{
		indexMap: map[int]int{},
		values:   []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.indexMap[val]
	if ok {
		return false
	}
	this.values = append(this.values, val)
	this.indexMap[val] = len(this.values) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// 最后一个数填到要删除的数的位置
	index, ok := this.indexMap[val]
	if !ok {
		return false
	}
	lastIndex := len(this.values) - 1
	lastVal := this.values[lastIndex]

	this.values[index] = lastVal
	this.indexMap[lastVal] = index

	// 删除元素
	this.values = this.values[:lastIndex]
	// todo map 和 数组都要删
	delete(this.indexMap, lastVal)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}
