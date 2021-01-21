package Interview

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var myDigits = []string{"", "",
		"abc", "def", "ghi",
	}

	// newList = append(newList, myDigits[myNumbers[0]:myNumbers[len(myNumbers)-1]+1]...)

	var toAppend []string
	mixMyString(myDigits, &toAppend, "", 0, digits)

	return toAppend
}

func mixMyString(myDigits []string, toAppend *[]string, currentStr string, index int, digits string) {
	if index == len(digits) {
		*toAppend = append(*toAppend, currentStr)
		return
	}

	// Catch one letter from myDigits, ex: a
	letters := myDigits[digits[index]-'0']
	fmt.Println(letters)
	// Loop through letters and " " +
	for i := 0; i < len(letters); i++ {
		fmt.Println(currentStr)
		mixMyString(myDigits, toAppend, currentStr+string(letters[i]), index+1, digits)
	}

}

func main() {
	x := letterCombinations("23")

	fmt.Println(x)
}

// Inside first loop
// "" + a - index = 1

// Inside second loop
// "a" + d - index = 2
// ["ad"]

// "a" + e - index = 2
// ["ae"]

// "a" + f - index =  2
// ["af"]

// return to first loop

// "" + b - index = 1
