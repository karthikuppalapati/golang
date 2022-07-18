package main

import "fmt"

//Example of a simple function without parameters and return values
func simpleGreet() {
	fmt.Println("Hello, how are you doing?")
}

//Example of a slightly modified function which takes a single parameter of string type
func greetWithName(name string) {
	fmt.Println("Hello", name)
}

//Example of a function with multiple parameters and a single return value
func simpleSum(a int, b int) int {
	res := a + b
	return res
	//We can do inline calculation without using a res variable  - return a + b
}

func multiSum(nums ...int) int { // The 3 dots tells the run time to take all the last unmatched ending int values and make them into a slice.
	var res int
	for _, v := range nums { //Here _ is used to store a value for which we dont want to create a variable.
		res += v
	}
	return res
}

//If there are multiple parameters of same primitive type then we can mention the type for all at once - Check parameters in line 16 and line 23.
//Example of a function with multiple parameters and return values
func divide(a, b float32) (float32, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Denominator can not be 0")
	}
	return a / b, nil
}

func changeValue(num *int) { // receives a pointer. Here num is a pointer type
	*num = 79 // * is used to dereference it
}

func main() {
	simpleGreet()                    // Prints "Hello, how are you doing?"
	greetWithName("Karthik")         // Prints "Hello Karthik"
	simpleSum(1, 2)                  // Returns an int value of 3. We need to store the returned value if we want to use it.
	multiSum(1, 2, 3, 4, 5, 6, 7, 8) //Returns an int value of 36
	a, err := divide(5, 3)           // Here a = 1.666 and err = nil
	if err != nil {
		fmt.Println(err)
	}
	b, err := divide(7, 0) // Here b = 0(as we assigned it consciously) and err = "Denominator can not be 0"
	if err != nil {
		fmt.Println(err) // Prints "Denominator can not be 0"
	}
	fmt.Println(a, b) // Prints 1.666 and 0.0

	//Anonymous functions
	func() {
		fmt.Println("Hello, I am anonymous function") // Prints "Hello, I am anonymous function" when invoked
	}() // These paranthesis call the function immediately

	//All the parameters passed to function calls are passed by value(original value does not change) except slices and maps
	num := 21
	changeValue(&num) // num is passed by reference hence its value may change
	fmt.Println(num)  //Prints 79
}
