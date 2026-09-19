package hash_map

import "math/rand"

/*

Implement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present(不存在). Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements
(it's guaranteed that at least one element exists when this method is called).
Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity.



Example 1:
Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output
[null, true, false, true, 2, true, false, 2]

Explanation
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.


Constraints:

-231 <= val <= 231 - 1
At most 2 * 105 calls will be made to insert, remove, and getRandom.
There will be at least one element in the data structure when getRandom is called.
*/

/*

Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output
[null, true, false, true, 2, true, false, 2]





*/

type RandomizedSet struct {
	// key=数值 value=切片的index
	indexMap map[int]int
	values   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indexMap: make(map[int]int),
		values:   make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.indexMap[val]
	if ok {
		return false
	}
	// todo append 到 slice，得到index，再保存到 map，否则会越界
	this.values = append(this.values, val)
	this.indexMap[val] = len(this.values) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	removeIndex, ok := this.indexMap[val]
	if !ok {
		return false
	}
	lastIndex := len(this.values) - 1
	lastVal := this.values[lastIndex]

	// todo 数组最后的元素移到要删除的元素的位置
	this.values[removeIndex] = lastVal
	// todo 移动的元素的index更新为删除元素的index
	this.indexMap[lastVal] = this.indexMap[val]
	this.values = this.values[:lastIndex]
	delete(this.indexMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
