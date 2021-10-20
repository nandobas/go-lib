package fizzbuzz

import (
	"strconv"
)

func Fizzbuzz(i int) string {
	var result string
	var	istr = strconv.Itoa(i)

	if( i%3 == 0) {
		result = "Fizz"
	}

	if( i%5 == 0) {
		result += "Buzz"
	}

	if( len(result) == 0) {
		result = istr
	}

	return result
}