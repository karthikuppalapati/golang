package main

import (
	"fmt"
	"reflect"
)

// Struct
/* type <structName> struct {
	field1 type1,
	field2 type2,
	...
} */
type teacher struct {
	name string
}
type student struct {
	name       string
	standard   int
	percentage float32
	guardian   string
	teacher    teacher // One struct nested inside of another struct
}

func main() {
	// Declaring a variable of struct type
	var student1 student
	fmt.Println(student1) // Here all are initialized to zero value. Prints - { 0  0  {}}

	// Struct Literal Syntax
	var student2 = student{name: "Tony", standard: 11, percentage: 77.62}
	fmt.Println(student2)                        // Prints - {Tony 11  77.62 {}}
	fmt.Println(reflect.TypeOf(student2).Kind()) // Prints - struct

	// Using the new keyword
	var student3 = new(student)                  // new() creates a pointer to the struct so student3 is a pointer type
	fmt.Println(reflect.TypeOf(student3).Kind()) // Prints - ptr

	// Accessing the struct fields using the . operator
	student3.name = "Patrick"
	student3.guardian = "Lisbon"
	student3.percentage = 99.9
	student3.teacher.name = "Bertram"
	fmt.Println(student3)  // Prints - &{Patrick 0  99.9 Lisbon {Bertram}}
	fmt.Println(*student3) // Prints - {Patrick 0  99.9 Lisbon {Bertram}}

	student4 := student2
	student4.percentage = 82        // Change of student4 fields did not effect student2 fields. So student4 is a copy of student2 but not pointed at studented2
	fmt.Println(student2, student4) // Prints - {Tony 11 77.62  {}} {Tony 11 82  {}}

	// Anonymous struct
	pizza := struct {
		name string
	}{
		name: "Pizza",
	}
	fmt.Println(pizza) // prints {Pizza}

}
