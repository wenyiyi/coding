package fixed_length_sliding_window

/*
https://www.hellointerview.com/learn/code/sliding-window/maximum-points-you-can-obtain-from-cards

Max Points You Can Obtain From Cards
DESCRIPTION (inspired by Leetcode.com)
Given an array of integers representing card values,
write a function to calculate the maximum score you can achieve by picking exactly k cards.

You must pick cards in order from either end. You can take some cards from the beginning,
then switch to taking cards from the end, but you cannot skip cards or pick from the middle.

For example, with k = 3:

Take the first 3 cards: valid
Take the last 3 cards: valid
Take the first card, then the last 2 cards: valid
Take the first 2 cards, then the last card: valid
Take card at index 0, skip some, then take card at index 5: not valid (skipping cards)
Constraints: 1 <= k <= cards.length

Example 1:
Input: cards = [2,11,4,5,3,9,2] k = 3
Output: 17

Explanation:
First 3 cards: 2 + 11 + 4 = 17
Last 3 cards: 3 + 9 + 2 = 14
First 1 + last 2: 2 + 9 + 2 = 13
First 2 + last 1: 2 + 11 + 2 = 15
Maximum score is 17.
*/

/*
Solution
Since you can only pick cards from the left or right, the cards you don't pick must stay together in the middle.
To get the max points = find the minimum remaining sum.
*/
func MaxScore(cards []int, k int) int {
	start := 0
	windowMax := 0
	windowK := len(cards) - k
	currentMin := 0
	totalScore := 0

	for end := range cards {
		totalScore += cards[end]
		windowMax += cards[end]
		if end < windowK {
			currentMin += cards[end]
		}

		if end-start+1 == windowK {
			currentMin = min(currentMin, windowMax)
			windowMax -= cards[start]
			start++
		}
	}
	return totalScore - currentMin
}
