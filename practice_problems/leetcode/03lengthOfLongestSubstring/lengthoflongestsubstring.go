package lengthoflongestsubstring

/*
################
PROBLEM
################

LeetCode difficulty: medium

Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/

/*
################
REFLECTION
################

I lost focus hard, and got caught up in a not-great first solution.

Counting distraction time, I spent about an hour on this one before taking a small
break. When I returned, I had the classic eureka moment of being able to see where
my approach was too convoluted and shitty, and was able to redesign my implementation
into a working solution within about ten minutes or so.
*/

/*
################
SOLUTION STATS
################

Runtime: 76 ms, faster than 22.36% of Go online submissions
Memory Usage: 2.6 MB, less than 94.55% of Go online submissions
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	contains := func(s string, c byte) bool {
		for i := 0; i < len(s); i++ {
			if c == s[i] {
				return true
			}
		}
		return false
	}

	longest := 1

	anchor, lead := 0, 1
	for lead < len(s) {
		// if next character is a duplicate of the current substring,
		// increment anchor and reset lead
		if contains(s[anchor:lead], s[lead]) {
			anchor++
			lead = anchor + 1
			continue
		}

		// otherwise, increment lead to include that new character
		// and update longest as appropriate
		lead++
		if diff := lead - anchor; diff > longest {
			longest = diff
		}
	}

	return longest
}
