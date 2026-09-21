package review

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

	pre  *Node
	next *Node
}

type LRUCache struct {
	capacity int
	cacheMap map[int]*Node // todo 我有点不知道要不要用引用，就是要 *Node，因为 map 和链表必须指向同一个 Node，如果存 Node value，会发生复制
	// values   []int todo 多余的不需要

	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.pre = head

	return LRUCache{
		capacity: capacity,
		cacheMap: make(map[int]*Node),
		head:     head, // todo 忘记赋予值了
		tail:     tail,
	}
}

// 不存在 head - A - B - tail        head - node - A - B - tail
// 存在  head - A - node - tail      head - A - tail        head -node- A - tail
func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cacheMap[key]
	if ok { // 存在
		// 先摘除
		this.Remove(node)
		// 修改值
		node.value = value
		// 移到head
		this.Add(node)
		return

	}
	// 不存在
	// 创建新节点 // todo 忘记赋予值了
	node = &Node{key: key, value: value}
	// 移到head
	this.Add(node)
	// 加入 map
	this.cacheMap[key] = node
	if len(this.cacheMap) > this.capacity {
		// 淘汰最久未使用的节点
		this.RemoveLRU()
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cacheMap[key]
	if !ok { // 不存在
		return -1
	}
	// todo 忘记要移动head了
	this.Remove(node)
	this.Add(node)
	return node.value
}

func (this *LRUCache) Remove(node *Node) {
	left := node.pre
	right := node.next

	// 左边连右边
	left.next = right
	// 右边连左边
	right.pre = left
}

// 不存在 head - A - B - tail        head - node - A - B - tail
func (this *LRUCache) Add(node *Node) {
	// todo 定义左右
	left := this.head
	right := this.head.next

	// node 自己连左右 todo 修改结构之前，先保存原来的节点
	node.pre = left
	node.next = right

	// 左右连自己
	left.next = node
	right.pre = node

	// 加入 map todo Add 只负责移到 head，因为已经存在的节点不需要再次添加map
	// this.cacheMap[node.key] = node
}

func (this *LRUCache) RemoveLRU() {
	// 淘汰最久未使用的节点
	node := this.tail.pre
	// 先摘除
	this.Remove(node)
	// 再删除
	delete(this.cacheMap, node.key)
}
