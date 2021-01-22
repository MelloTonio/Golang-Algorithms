package main

import (
	"fmt"
	"strings"
)

func largestCommonPrefix(strarr []string) string {
	if len(strarr) == 0 {
		return ""
	}

	myPrefix := strarr[0]

	for i := 1; i < len(strarr); i++ {
		fmt.Println(strings.Contains(strarr[i], myPrefix))
		for strings.HasPrefix(strarr[i], myPrefix) != true {
			//fmt.Println(strings.Contains("flight", "fl"))
			myPrefix = myPrefix[:len(myPrefix)-1]

		}
	}

	return myPrefix
}
