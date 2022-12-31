package main

import (
	"errors"
	"fmt"
)

func main() {
	num, den, err := divAndRemainder(5, 2)
	fmt.Println(num, den, err)
}

func divAndRemainder(numerator, denominator int) (result int, remainder int,
	err error) {
	// assign some values
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
