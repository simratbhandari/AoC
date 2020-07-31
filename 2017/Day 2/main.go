package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), "\t")
		A := make([]int, len(words))
		for i := 0; i < len(words); i++ {
			v, err := strconv.Atoi(words[i])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error parsing integer '%v': %v\n", words[i], err)
				os.Exit(1)
			}
			A[i] = v
		}
		sum += findEvenQuotient(A)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", sum)
}

// Part 1
// The function returns in O(n) time.
func computeRange(A []int) int {
	min := A[0]
	max := A[0]
	for i := 1; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}
	return max - min
}

// Part 2
// O(n^2)
func findEvenQuotient(A []int) int {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			a := A[i]
			b := A[j]
			if a < b {
				a, b = b, a
			}
			if a%b == 0 {
				return a / b
			}
		}
	}
	panic("this should never happen")
}
