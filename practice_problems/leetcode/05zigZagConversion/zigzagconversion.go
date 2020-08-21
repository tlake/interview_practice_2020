package zigzagconversion

/*
################
PROBLEM
################

LeetCode difficulty: medium

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I

*/

/*
################
REFLECTION
################

Time: 20 min to come up with approach and get broad strokes of code written.
Another 10 to bugfix, which included off-by-one errors in the Oscillator,
bad syntax while trying to create a slice of byte-slices, and type incompatibilities.
Those are the kinds of mistakes that ngrok is gonna get super tetchy about.
*/

/*
################
SOLUTION STATS
################

Time complexity: O(n)
Space complexity: O(n)

Runtime: 20 ms, faster than 25.00% of Go online submissions for ZigZag Conversion.
Memory Usage: 6.2 MB, less than 52.62% of Go online submissions for ZigZag Conversion.
*/

type Oscillator struct {
	Range   []int
	Current int // index
}

func (o *Oscillator) Next() int {
	curr := o.Range[o.Current]

	if o.Current == len(o.Range)-1 {
		o.Current = 0
	}
	o.Current++

	return curr
}

func NewOscillator(numItems int) *Oscillator {
	o := &Oscillator{}

	for i := 0; i < numItems; i++ {
		o.Range = append(o.Range, i)
	}

	for i := numItems - 2; i >= 0; i-- {
		o.Range = append(o.Range, i)
	}

	return o
}

func convert(s string, numRows int) string {
	// early edgecase exit
	if numRows == 1 {
		return s
	}

	o := NewOscillator(numRows)

	rows := make([][]byte, numRows)

	// append each char in s to a row in rows as determined by the oscillator
	for i := 0; i < len(s); i++ {
		rowIndex := o.Next()
		rows[rowIndex] = append(rows[rowIndex], s[i])
	}

	// concatenate the rows
	for i := 1; i < len(rows); i++ {
		rows[0] = append(rows[0], rows[i]...)
	}

	return string(rows[0])
}
