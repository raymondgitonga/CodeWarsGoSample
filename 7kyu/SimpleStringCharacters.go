package _kyu

import (
	"strings"
	"unicode"
)

// Source : https://www.codewars.com/kata/5a29a0898f27f2d9c9000058/

/**********************************************************************************
 *
 * In this Kata, you will be given a string and your task will be to return a list of ints detailing
 * the count of uppercase letters, lowercase, numbers and special characters, as follows.
 *
 * Solve("*'&ABCDabcde12345") = [4,5,5,3].
 * --the order is: uppercase letters, lowercase, numbers and special characters.
 **********************************************************************************/

func SolveK(s string) []int {

	alphU := 0
	alphL := 0
	num := 0
	char := 0

	for _, r := range s {
		if unicode.IsLetter(r) {
			if strings.ToUpper(string(r)) == string(r) {
				alphU += 1
			} else {
				alphL += 1
			}
		} else if unicode.IsNumber(r) {
			num += 1
		} else {
			char += 1
		}
	}

	return []int{alphU, alphL, num, char}
}
