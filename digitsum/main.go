package main

import (
	"fmt"
	"os"
	"strconv"
)

// computeDigitSum calculates the digit sum of an integer number
func computeDigitSum(number int) int {
	/*
		TODO implement computation of the sum of all digits

		Examples:
			7 => 7
			21 => 3
			101 => 2

		Bonus:
		Once you are done, you might try to extend this code to support huge numbers, by
		using big.Int as input parameter instead of int.
	*/
	return -1
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Missing required number argument.")
		os.Exit(1)
	}

	numParam := os.Args[1]
	n, err := strconv.Atoi(numParam)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parameter '%v' is not a number: %v\n", numParam, err)
		os.Exit(2)
	}

	dsum := computeDigitSum(n)
	fmt.Println(dsum)
}
