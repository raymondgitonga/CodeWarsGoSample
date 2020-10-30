package _kyu

// Source : https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9

/**********************************************************************************
 *
 * Simple, given a string of words, return the length of the shortest word(s).
 * String will never be empty and you do not need to account for different data types.
 *
 **********************************************************************************/

import "strings"

func FindShort(s string) int {
	wordArr := strings.Split(s, " ")
	smallest := len(wordArr[0])

	for i := 1; i < len(wordArr); i++ {
		if len(wordArr[i]) < smallest {
			smallest = len(wordArr[i])
		}
	}
	return smallest
}
