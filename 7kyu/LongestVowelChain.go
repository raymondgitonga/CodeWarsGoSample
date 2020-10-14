package _kyu

// Source : https://www.codewars.com/kata/59c5f4e9d751df43cf000035

/**********************************************************************************
 *
 * The vowel substrings in the word codewarriors are o,e,a,io. The longest of these has a length of 2.
 * Given a lowercase string that has alphabetic characters only (both vowels and consonants) and no spaces,
 * return the length of the longest vowel substring. Vowels are any of aeiou.
 *
 **********************************************************************************/
import "regexp"

func Solve(s string) int {

	re := regexp.MustCompile(`[a e i o u]+`)
	results := re.FindAllString(s, -1)
	max := 0
	for _, s := range results {
		if max < len(s) {
			max = len(s)
		}
	}
	return max
}