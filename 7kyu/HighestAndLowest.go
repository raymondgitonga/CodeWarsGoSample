package _kyu

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Source : https://www.codewars.com/kata/554b4ac871d6813a03000035

/**********************************************************************************
 *
 * In this little assignment you are given a string of space separated numbers,
 * and have to return the highest and lowest number.
 *
 * highAndLow("1 2 3 4 5");  // return "5 1"
 * highAndLow("1 2 -3 4 5"); // return "5 -3"
 * highAndLow("1 9 3 4 -5"); // return "9 -5"
 **********************************************************************************/

func HighAndLow(in string) string {
	arrStr := strings.Split(in, " ")
	var arrInt []int

	for _, in := range arrStr {
		convInt, _ := strconv.Atoi(in)
		arrInt = append(arrInt, convInt)
	}
	sort.Ints(arrInt)

	return fmt.Sprintf("%v %v", arrInt[len(arrInt)-1], arrInt[0])
}
