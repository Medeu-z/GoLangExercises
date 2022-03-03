// Homework 0: Hello Go!

package main

import (
	"fmt"
)

func main() {
	// Feel free to use the main function for testing your functions
	//Testing Fizzbuzz function
	fmt.Println("Testing Fizzbuzz function")
	fmt.Println("  6: ", Fizzbuzz(6))
	fmt.Println(" 10: ", Fizzbuzz(10))
	fmt.Println(" 15: ", Fizzbuzz(15))
	fmt.Println(" 23: ", Fizzbuzz(23))

	//Testing IsPrime function
	fmt.Println("\nTesting IsPrime function")
	fmt.Println("Is  1 is prime number: ", IsPrime(1))
	fmt.Println("Is 27 is prime number: ", IsPrime(27))
	fmt.Println("Is 29 is prime number: ", IsPrime(29))

	//Testing ISPalindrome function
	fmt.Println("\nTesting IsPalindrome function")
	fmt.Println("Is kazak palindrome: ", IsPalindrome("kazak"))
	fmt.Println(" Is toot palindrome: ", IsPalindrome("toot"))
	fmt.Println(" Is this palindrome: ", IsPalindrome("this"))
	fmt.Println("  Is kaz palindrome: ", IsPalindrome("kaz"))

}

// Fizzbuzz is a classic itrductory programming problem.
// If n s divisibleby 3, return "Fizz"
// If n is divisible y 5,return "Buzz"
// If n is diviibl by 3 and 5, return "FizzBuzz"
// Otherwise, return e empty string
func Fizzbuzz(n int) string {
	// TODO
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%5 == 0 {
		return "Buzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else {
		return ""
	}

}

// IsPrime cheks if the number is prime.
//You may use any prime agorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// TODO
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}

// sPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	// TODO
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}

	}
	return true
}
