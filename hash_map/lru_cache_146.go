package hash_map

/*
https://leetcode.com/problems/lru-cache/description/

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists.
Otherwise, add the key-value pair to the cache.
If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.


Example 1:
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4


Constraints:
1 <= capacity <= 3000
0 <= key <= 104
0 <= value <= 105
At most 2 * 105 calls will be made to get and put.

*/

/*
Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4

map 1:1,2:2

head是最新使用的:{2,2}-{1,1}
get{1}
{1,1}-{2,2}
put(3, 3)
{3,3}-{1,1}-{2,2}
移除{2,2}





*/

type Node struct {
	key   int
	value int
	next  *Node
	pre   *Node
}

type LRUCache struct {
	capacity int
	// hashmap todo key → Node
	cacheMap map[int]*Node
	// todo doubleLinkList
	head *Node
	tail *Node
}

func ConstructorLRUCache(capacity int) LRUCache {
	// todo head 和 tail 需要初始化
	head := &Node{}
	tail := &Node{}

	// todo head ⇄ tail
	head.next = tail
	tail.pre = head

	return LRUCache{
		capacity: capacity,
		cacheMap: make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	// 从原来的位置摘下来
	this.Remove(node)
	// 移到head后面
	this.Add(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cacheMap[key]
	if ok { // 找到node
		// 更新值
		node.value = value
		// 从原来的位置摘下来
		this.Remove(node)
		// 移到head后面
		this.Add(node)
		return
	}
	// node不存在，为nil，需要新创建node
	node = &Node{}
	node.key = key
	node.value = value
	this.cacheMap[key] = node
	// 添加到head
	this.Add(node)
	if len(this.cacheMap) > this.capacity {
		// 删除尾元素
		this.Delete()
	}

}

// 把节点摘下来
// left := node.pre
// right := node.next
// A ⇄ X ⇄ B  删除X
// A ⇄ B todo 删除后，左右牵手
func (this *LRUCache) Remove(node *Node) {
	left := node.pre   // A
	right := node.next // B

	// A 连 B，B 连 A
	left.next = right
	right.pre = left
}

// 把节点移到head后面
// A ⇄ B todo 插入后，自己连左右
// A ⇄ X ⇄ B
func (this *LRUCache) Add(node *Node) {
	// 定义左右是谁
	left := this.head
	right := this.head.next

	// 节点自己连左右
	node.pre = left
	node.next = right
	left.next = node
	right.pre = node
}

// head -> 1->2 -> tail
func (this *LRUCache) Delete() {
	// 需要删除
	node := this.tail.pre
	// todo 直接执行摘除
	this.Remove(node)
	// todo 从map删除
	delete(this.cacheMap, node.key)
}
