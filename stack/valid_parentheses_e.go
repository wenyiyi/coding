package stack

/*
https://www.hellointerview.com/learn/code/stack/valid-parentheses

Valid Parentheses
DESCRIPTION (inspired by Leetcode.com)
Given an input string s consisting solely of the characters '(', ')', '{', '}', '[' and ']', determine whether s is a valid string. A string is considered valid if every opening bracket is closed by a matching type of bracket and in the correct order, and every closing bracket has a corresponding opening bracket of the same type.

Example 1:
Inputs: s = "(){({})}"
Output: True

Example 2:
Inputs: s = "(){({}})"
Output: False
*/
func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, bracket := range s {
		switch bracket {
		case '(', '{', '[':
			// push
			stack = append(stack, bracket)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[bracket] {
				return false
			}
			// stack = stack[:len(stack)-1] pop: delete the top element
			// stack = stack[len(stack)-1]  peek: get top element
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}
