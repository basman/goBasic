package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"calculator/calc"
)

// simple calculator

func prompt(scanner *bufio.Scanner) (string, error) {
	fmt.Printf("? ")

	if !scanner.Scan() {
		return "", nil
	}

	line := scanner.Text()

	if line == "" {
		return "", nil
	}

	if line == "q" {
		return "", nil
	}

	return line, nil
}

func compute(in string) (string, error) {
	out, err := calc.Eval(in)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", out), nil
}

func main() {
	if len(os.Args) <= 1 {
		runStdin()
		return
	}

	out, err := compute(strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(out)
}

func runStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		in, err := prompt(scanner)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
			continue
		}

		if in == "" {
			break
		}

		out, err := compute(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Computation error: %v\n", err)
			continue
		}

		fmt.Println(out)
	}
}
