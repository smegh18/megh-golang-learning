package main

import (
	"fmt"
	"math"
)

// Check if a number is odd
func odd(n int) bool {
	return n%2 != 0
}

// Check if a number is even
func even(n int) bool {
	return n%2 == 0
}

// Check if a number is prime
func prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Check if a number is greater than a given value
func greaterThan(n, threshold int) bool {
	return n > threshold
}

// Check if a number is a multiple of a given number
func multipleOf(n, divisor int) bool {
	return n%divisor == 0
}

// Story 1: Return even numbers
func filterEven(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if even(num) {
			result = append(result, num)
		}
	}
	return result
}

// Story 2: Return odd numbers
func filterOdd(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if odd(num) {
			result = append(result, num)
		}
	}
	return result
}

// Story 3: Return prime numbers
func filterPrime(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if prime(num) {
			result = append(result, num)
		}
	}
	return result
}

// Story 4: Return odd prime numbers
func filterOddPrime(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if prime(num) && odd(num) {
			result = append(result, num)
		}
	}
	return result
}

// Story 5: Return even numbers that are multiples of 5
func filterEvenMultiplesOf5(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if even(num) && multipleOf(num, 5) {
			result = append(result, num)
		}
	}
	return result
}

// Story 6: Return odd numbers that are multiples of 3 and greater than 10
func filterOddMultiplesOf3GreaterThan10(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if odd(num) && multipleOf(num, 3) && greaterThan(num, 10) {
			result = append(result, num)
		}
	}
	return result
}

// Story 7: Return numbers that match all conditions (AND logic)
func filterAll[T any](numbers []T, conditions ...func(T) bool) []T {
	var result []T
	for _, num := range numbers {
		matches := true
		for _, cond := range conditions {
			if !cond(num) {
				matches = false
				break
			}
		}
		if matches {
			result = append(result, num)
		}
	}
	return result
}

// Story 8: Return numbers that match any condition (OR logic)
func filterAny[T any](numbers []T, conditions ...func(T) bool) []T {
	var result []T
	for _, num := range numbers {
		for _, cond := range conditions {
			if cond(num) {
				result = append(result, num)
				break
			}
		}
	}
	return result
}

// Test Functions (for testing and validation)
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Story 1
	fmt.Println("Even Numbers:", filterEven(numbers))
	// Story 2
	fmt.Println("Odd Numbers:", filterOdd(numbers))
	// Story 3
	fmt.Println("Prime Numbers:", filterPrime(numbers))
	// Story 4
	fmt.Println("Odd Prime Numbers:", filterOddPrime(numbers))
	// Story 5
	fmt.Println("Even Multiples of 5:", filterEvenMultiplesOf5(numbers))
	// Story 6
	fmt.Println("Odd Multiples of 3 Greater Than 10:", filterOddMultiplesOf3GreaterThan10(numbers))

	// Story 7 (All conditions)
	allConditions := []func(int) bool{
		odd,
		func(n int) bool { return greaterThan(n, 5) },
		func(n int) bool { return multipleOf(n, 3) },
	}
	fmt.Println("All Conditions (Odd, Greater Than 5, Multiple of 3):", filterAll(numbers, allConditions...))

	// Story 8 (Any condition)
	anyConditions := []func(int) bool{
		prime,
		func(n int) bool { return greaterThan(n, 15) },
		func(n int) bool { return multipleOf(n, 5) },
	}
	fmt.Println("Any Condition (Prime, Greater Than 15, Multiple of 5):", filterAny(numbers, anyConditions...))
}
