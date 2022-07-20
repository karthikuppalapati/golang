package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array declaration without initialization
	var nums [3]int
	fmt.Println(nums) // Prints [0,0,0] because if not initialized, default value will be 0 for int data type

	// Initializing array inline
	var names [4]string = [4]string{"Iron man", "Thor", "Hulk", "Captain Marvel"}
	fmt.Println(names) // Prints [Iron man Thor Hulk Captain Marvel]

	// Shorthand initializing
	userNames := [2]string{"coolDude", "i_am_rocky"}
	fmt.Println(userNames) // Prints [coolDude i_am_rocky]

	// We can initialize/update/access array values using the indices(starts at 0)
	nums[0], nums[2] = 26, 45
	fmt.Println(nums, nums[2]) // Prints [26 0 45] 45

	var marks [10]int // First parameter is the type and second parameter is the size

	for i, v := range marks { // Here i is the index and v is the value at that index i.e v = marks[i]
		fmt.Println(i, v)
		marks[i] = i + 1
	}
	fmt.Println(marks) // Prints [1 2 3 4 5 6 7 8 9 10]

	// len() function gives the length of the array
	fmt.Println(len(marks)) // Prints 10

	fmt.Println(reflect.TypeOf(marks).Kind()) // Gives the type of the variable - In this case it prints "array"

	// Multi dimensional array
	var matrix [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	} // Prints - 1 2
	//		  3 4

}
