package stack

import "strings"

/*
https://www.hellointerview.com/learn/code/stack/decode-string

Decode String
DESCRIPTION (inspired by Leetcode.com)
Given an encoded string s, write a function to return its decoded string.

The encoding rule is k[encoded_string], where the encoded_string inside the square brackets is repeated exactly k times.
k is always a positive integer, and the brackets can be nested.

You can assume the input is always well-formed: there are no extra spaces, every square bracket is properly matched,
and digits only ever appear to specify a repeat count k (so you won't see input like 3a or 2[4]).

Constraints:
1 <= s.length
s consists of lowercase English letters, digits, and the square brackets [ and ].
All repeat counts k are positive integers and may have more than one digit (for example, 10[a]).
The input string is guaranteed to be valid.
Example 1:
Inputs: s = "3[a]2[bc]"
Output: "aaabcbc"
(Explanation: 3[a] decodes to "aaa" and 2[bc] decodes to "bcbc".)

Example 2:
Inputs: s = "3[a2[c]]"
Output: "accaccacc"
(Explanation: the inner 2[c] becomes "cc", so a2[c] is "acc", which is then repeated 3 times.)

Example 3:
Inputs: s = "efg2[abc]3[cd]ef"
Output: "efgabcabccdcdcdef"
*/
func decodeString(s string) string {
	currString := ""
	currNumber := 0
	var stringStack []string
	var numberStack []int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			// It may be a multi-digit number. 有可能为多位数
			currNumber = currNumber*10 + int(c-'0')
		} else if c == '[' { // 遇到 [ 就记住前面的东西,记住当前数字和当前字符串
			stringStack = append(stringStack, currString) // ["efg"]
			numberStack = append(numberStack, currNumber) // 2
			currString = ""
			currNumber = 0
		} else if c == ']' { // 遇到 ] 就把里面的东西重复,然后拼回去
			// get top element
			num := numberStack[len(stringStack)-1]
			// delete the top element
			numberStack = numberStack[:len(numberStack)-1]
			prevString := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]
			currString = prevString + strings.Repeat(currString, num)
		} else {
			currString += string(c)
		}

	}

	return currString
}
