package Interview

func BalancedParenthesis(n int) []string {
	var myArray = []string{}

	backtrack(&myArray, "", 0, 0, n)

	return myArray

}

func backtrack(arr *[]string, parenthesis string, open int, close int, max int) {
	// Base case - If we have all the needed parenthesis we can return
	if len(parenthesis) == max*2 {
		*arr = append(*arr, parenthesis)
		return
	}

	if open < max {
		backtrack(arr, parenthesis+"(", open+1, close, max)
	}

	if close < open {
		backtrack(arr, parenthesis+")", open, close+1, max)
	}

}
