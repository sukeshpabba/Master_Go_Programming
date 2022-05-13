/////////////////////////////////
// The Typical Structure of a Go Application
// Go Playground: https://play.golang.org/p/vY_IeYBb1GN
/////////////////////////////////

// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main

// import declaration declares packages used in this file
import (
	"fmt"
)

// package scoped variables and constants. you cannot use := in package scoped variables & constants
var x int = 100

const y = 0

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main() {

	// Local Scoped Variables and Constants Declarations, calling functions etc
	//Actual Declaration of Variables using "var" keyword
	var a int = 7
	var b float64 = 3.5
	const c int = 10
	//Short Declaration Operator
	d := 123
	e := "Short Declaration Operator"
	f := 1.2
	fmt.Println(x)
	fmt.Println(y)
	// Println() function prints out a line to stdout
	// It belongs to package fmt
	fmt.Println("Hello Go world!") // => Hello Go world!
	fmt.Println(a, b, c, d, e, f)  // => 7 3.5 10
	//** DECLARING VARIABLES **///

	// Syntax: var var_name type
	var s11 string = "Learning Go!"
	fmt.Println(s11) // printing string s1

	//** TYPE INFERENCE **//

	// Go deduces automatically the type of the variable by looking at the initial value (bool, int, string etc)

	var k int = 6 // not necessary to say the type (int). It is inferred from the literal on the right side of =
	var i = 5     // type int
	var j = 5.6   // type float64

	// printing i, j and k
	fmt.Println("i:", i, "j:", j, "k:", k)

	// ii == jj  // -> error: cannot assign float to int (Go is a strong typed language)

	// declaring and initializing a new variable of type string (type inference)
	var s2 = "Go!"
	_ = s2 //in Go each variable must be used or there is a compile-time error
	// _ is the Blank Identifier and mutes the error of unused variables
	// _ can be only on the left hand side of the = operator

	// multiple assignments
	var ii, jj int
	ii, jj = 5, 8 // -> tuple assignment. It allows several variables to be assigned at once

	// swapping two variables using multiple assignments
	ii, jj = jj, ii

	fmt.Println(ii, jj)

	//** Short Declaration (works only in Block Scope!) **//

	// := (colon equals syntax) used only when declaring a new variable (or at least a new variable)
	// := tells go we are going to create a new variable and go figures out what type it will be
	s3 := "Learning golang!"
	_ = s3

	// can't use short declaration at Package Scope (outside main() or other function)
	// all statements at package scope must start with a Go keyword (package, var, import, func etc)

	// multiple short declaration
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	// redeclaration with short declaration syntax
	// at least one variable must be NEW on the left side of :=
	var deleted = false
	deleted, file := true, "a.txt"
	_, _ = deleted, file

	// expressions in short declarations are allowed
	sum := 5 + 2.5
	fmt.Println(sum)

	// multiple declaration is good for readability
	var (
		age       float64
		firstName string
		gender    bool
	)
	_, _, _ = age, firstName, gender

	// a concise way to declare multiple variables that have the same type
	var a1, b1 int
	_, _ = a1, b1

}
