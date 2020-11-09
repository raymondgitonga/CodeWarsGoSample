package _kyu

import (
	"strings"
)

// Source : https://www.codewars.com/kata/5b180e9fedaa564a7000009a

/**********************************************************************************
 *
 * In this Kata, you will be given a string that may have mixed uppercase and lowercase letters and your
 * task is to convert that string to either lowercase only or uppercase only based on:
 * make as few changes as possible.
 * if the string contains equal number of uppercase and lowercase letters, convert the string to lowercase.
 * For example:
 * solve("coDe") = "code". Lowercase characters > uppercase. Change only the "D" to lowercase.
 * solve("CODe") = "CODE". Uppercase characters > lowercase. Change only the "e" to uppercase.
 * solve("coDE") = "code". Upper == lowercase. Change all to lowercase.
 **********************************************************************************/

func solve(str string) string {
	var upper int
	var lower int

	strArr := strings.Split(str, "")

	for _, st := range strArr {
		if strings.ToLower(st) == st {
			lower += 1
		} else {
			upper += 1
		}
	}

	if lower == upper || lower > upper {
		return strings.ToLower(str)
	} else {
		return strings.ToUpper(str)
	}

}
