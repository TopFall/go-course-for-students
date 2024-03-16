package main

import (
	"fizzbuzz"
	"fmt"
)

func main() {
	// TODO тут напишите цикл с вызовом FizzBuzz
	for i := 0; i <= 100; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
