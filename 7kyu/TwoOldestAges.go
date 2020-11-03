package _kyu

import "sort"

// Source :https://www.codewars.com/kata/511f11d355fe575d2c000001

/**********************************************************************************
 *
 * The two oldest ages function/method needs to be completed. It should take an array of numbers as its argument and
 * return the two highest numbers within the array. The returned value should be an array in the format
 * [second oldest age, oldest age].

 * The order of the numbers passed in could be any order. The array will always include at least 2 items.
 * For example:
 * TwoOldestAges([]int{1, 5, 87, 45, 8, 8}) // should return [2]int{45, 87}
 **********************************************************************************/

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)

	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}
