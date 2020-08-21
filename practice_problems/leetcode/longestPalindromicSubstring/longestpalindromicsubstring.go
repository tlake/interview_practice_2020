package longestpalindromicsubstring

/*
################
PROBLEM
################

LeetCode difficulty: medium

Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd"
Output: "bb"
*/

/*
################
REFLECTION
################

Time: ~40 min

Still took way too long, still got stuck trying to implement a bad solution for too long. Still
got distracted. Maybe it's not a good environment today, but I keep getting distracted by fuckin'
nothing.

I need to improve my process, maybe try to talk my way through a solution on a whiteboard
before getting to code? Not sure if that's a good way to practice, since interviews done
during pandemic times will likely be entirely situated around virtual shared IDEs...
*/

/*
################
SOLUTION STATS
################

Runtime complexity: O(n^2) -- the isPalindrome function is O(log[2] n), but that'll be eclipsed
by the O(n^2) nested for-loop.

Space complexity: O(n) -- the substr and logestPalindrome vars make copy parts of the input,
potentially including the entire input. this could be refactored to use bracket notation to
access the input for substr, and longestPalindrome could be replaced with tracking the longest
palindrome's starting index and length to make most of program consume O(1), but isPalindrome
would still make copies of chunks of the input for its own input when called, though.

Runtime: 196 ms, faster than 22.43% of Go online submissions
Memory Usage: 2.3 MB, less than 59.81% of Go online submissions
*/

func longestPalindrome(s string) string {
	// quick edgecase exit
	if len(s) <= 1 {
		return s
	}

	// O(log[2] n)
	isPalindrome := func(s string) bool {
		l, r := 0, len(s)-1
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	longestPalindrome := s[:1]

	l, r := 0, len(s)
	// if the unchecked portion of the original string is smaller than the longest palindrome
	// found so far, there's no need to keep checking
	//
	// worst case, s contains only single-char palindromes, so we have to check every
	// item
	// --> O(n)
	for (r - l) > len(longestPalindrome) {
		comparator := r
		// check every substring that starts at l for palindromity (largest first so we can
		// break once we find one)
		//
		// worse case, s contains only single-char palindromes, so we have to compare each
		// remaining element to the right of l
		// --> O(n) again; this makes O(n^2) being nested like it is
		for comparator > l {
			substr := s[l:comparator]
			// if we find a palindrome, we can break out of the loop since any subsequent
			// palindromes starting at l would be smaller than the one we've just found
			if isPalindrome(substr) == true {
				if len(substr) > len(longestPalindrome) {
					longestPalindrome = substr
					break
				}
			}
			// otherwise, keep looking for the next-biggest potential palindrome
			comparator--
		}

		// move on to check all the substrings that start at the next l
		l++
		r = len(s)
	}

	return longestPalindrome
}
