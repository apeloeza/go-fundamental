package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		scores := []int{10,5,8,9,7}
		total := sum(scores)
	*/
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	/*
		result, err := calculate(10, 2, "+")
		result, err := calculate(10, 2, "-")
		result, err := calculate(10, 2, "*")
		result, err := calculate(10, 2, "=")
	*/
	result, err := calculate(10, 2, "s")
	//kita cek error
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("Unknown operation")
	}

	return result, errorResult
}

func sum(numbers []int) int {
	// kita melakukan looping
	var total int
	for _, number := range numbers {
		total = total + number
	}

	return total
}
