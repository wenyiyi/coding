package hash_map

/*
Given an array of strings strs, group the anagrams together.
You can return the answer in any order.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:
There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Constraints:
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

/*
需要让 eat、tea、ate 自动进入同一组，它们需要有同样的key
如果简单累加
hash = int(e) + int(a) + int(t)，容易出现hash collision
如：
"ad"
"bc"
a + d = 97 + 100 = 197
b + c = 98 + 99  = 197

小写字母和数字的映射
a → 0
b → 1
c → 2
...
z → 25

index := char - 'a'

0 1 2 3 4 5 6 7 8 9 10
a b c d e f g h i j k

1       1
eat
tea
*/
func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	// groups:
	// frequency1 → ["eat", "tea", "ate"]
	// frequency2 → ["tan", "nat"]
	// frequency3 → ["bat"]
	for _, str := range strs {
		key := [26]int{}
		// 注意是 _,ch 而不是 直接 ch
		for _, ch := range str {
			// 每个字母的下标 index := char - 'a'
			key[ch-'a']++
		}
		groups[key] = append(groups[key], str)
	}

	// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	var result [][]string
	// 把 groups 的 value 放入result，
	for _, v := range groups {
		// 对于长度为 0 的 slice，想不断往里面增加元素，应该用 append
		result = append(result, v)
	}
	return result
}

/*
假设有 n 个字符串，每个字符串平均长度是 k

Time Complexity: O(n × k)
Space Complexity: O(n * k)，n个字符串，每个字符串平均长度 k，所有字符串总大小 ≈ n × k


n = 6       // 6 个字符串
k ≈ 3       // 每个字符串平均 3 个字符
strs := []string{
    "eat",
    "tea",
    "tan",
    "ate",
    "nat",
    "bat",
}
                         groups
                           │
          ┌────────────────┼────────────────┐
          │                │                │
          ▼                ▼                ▼

 key: [26]int         key: [26]int      key: [26]int]
  a=1,e=1,t=1         a=1,n=1,t=1       a=1,b=1,t=1
          │                │                │
          ▼                ▼                ▼

     []string           []string          []string
   ┌───────────┐       ┌─────────┐        ┌─────┐
   │ "eat"     │       │ "tan"   │        │"bat"│
   │ "tea"     │       │ "nat"   │        └─────┘
   │ "ate"     │       └─────────┘
   └───────────┘

     group 1             group 2           group 3

每个key [26]int 永远26个数字
6个字符串 * 3个字符 = n * k

*/
