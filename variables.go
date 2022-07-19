package main

import "fmt"

func main() {
	//There are various ways to declare the variables in Go
	// Naming convention - camelCase. If first letter of variable is Capital then it gets exported
	// Variables with no initialization
	var Name string // Gets exported because of Capital N at the start.
	var istrue bool
	/* If a variable is not initialized then it should have its data type and
	the variable has a default value of
		0 if int,
		false if bool
		"" if string
	*/
	// Variables with initialization
	var num = 10        // No type declaration
	var number int = 42 // Type declaration
	// We can not have two variables with same name. If we declare a variable we must use it
	fmt.Println(Name, num, number, istrue) // Prints " 10 42 false"

	// Shorthand syntax
	username := "karthik"
	myNumber := 54
	fmt.Println(username, myNumber) // prints "karthik 54"

	// You can initialize multiple variables at once by separating them with commas
	studentName1, teacherName, highestMarks := "Harry", "Dumbledore", 53
	println(studentName1, teacherName, highestMarks) // Prints "Harry Dumbledore 53"

	/* Unlike other languages, Go has many data types
	Integer: int8, int16, int32, int63, uint8, uint16, uint32, uint64
	Float:  float32, float64, complex64, complex128 (Go also has complex numbers)
	*/

	// Strings are immutable in Go
}
