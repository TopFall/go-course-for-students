package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else if i%15 == 0 {
		return "FizzBuzz"
	}
	return strconv.Itoa(i)
}
