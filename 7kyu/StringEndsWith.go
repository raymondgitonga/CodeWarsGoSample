package _kyu

// Source : https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d

/**********************************************************************************
 *
 * Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).
 * Examples:
 * solution("abc", "bc") // returns true
 * solution("abc", "d") // returns false
 **********************************************************************************/

func solution(str, ending string) bool {
	lenOrg := len(str)
	lenEnd := len(ending)

	if len(ending) > len(str) {
		return false
	}
	substr := str[lenOrg-lenEnd:]

	return ending == substr

}
