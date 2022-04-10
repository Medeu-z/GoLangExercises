// Homework 3: Interfaces
// Due February 14, 2017 at 11:59pm
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	hello := map[string]string{
		"hello":   "world",
		"hola":    "mundo",
		"bonjour": "monde",
	}
	for k, v := range hello {
		fmt.Printf("%s, %s\n", strings.Title(k), v)
	}
}

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces!
// To sort a collection (such as a slice), the type must satisfy sort.Interface,
// which requires 3 methods: Len() int, Less(i, j int) bool, and Swap(i, j int).
// To actually sort a slice, you need to first implement all 3 methods on a
// custom type, and then call sort.Sort on your the PersonSlice type.
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type PersonSlice []*Person

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.

var count int = 1

func NewPerson(first, last string) *Person {
	p := new(Person)
	p.ID = count
	p.FirstName = first
	p.LastName = last
	count++
	return p
}

func (ps PersonSlice) Less(i, j int) bool {
	if c := strings.Compare(ps[i].LastName, ps[j].LastName); c != 0 {
		return c < 0
	}
	if c := strings.Compare(ps[i].FirstName, ps[j].FirstName); c != 0 {
		return c < 0
	}
	return ps[i].ID < ps[j].ID
}

func (ps PersonSlice) Len() int {
	return len(ps)
}

func (ps PersonSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func IsPalindrome(s sort.Interface) bool {

	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, s.Len()-1-i) || s.Less(s.Len()-1-i, i) {
			return false
		}
	}

	return true
}

// Problem 3: Functional Programming
// Write a function Fold which applies a function repeatedly on a slice,
// producing a single value via repeated application of an input function.
// The behavior of Fold should be as follows:
//   - When s is empty, return v (default value)
//   - When s has 1 value (x0), apply f once: f(v, x0)
//   - When s has 2 values (x0, x1), apply f twice, from left to right: f(f(v, x0), x1)
//   - Continue this pattern recursively to obtain the final result.

// Fold applies a left to right application of f on s starting with v.
// Note the argument signature of f - func(int, int) int.
// This means f is a function which has 2 int arguments and returns an int.
func Fold(s []int, v int, f func(int, int) int) int {
	if len(s) == 0 {
		return v
	}

	return Fold(s[1:], f(v, s[0]), f)
}
