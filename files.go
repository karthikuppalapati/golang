package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	text := "Hey there how are you doing."
	file, err := os.Create("./newFile.text")
	checkNilError(err)

	writeIntoFile(file, text)
	readFile("./newFile.text")
	defer file.Close()
}
func writeIntoFile(file *os.File, data string) {
	length, err := file.WriteString(data) // Writing data into the file
	checkNilError(err)
	fmt.Println(length, " bytes were written into the file")
}

func readFile(fileName string) {
	databytes, err := ioutil.ReadFile(fileName) // databytes is a slice of UTF-8 form of characters read
	checkNilError(err)
	fmt.Println("Data in the file is : ", string(databytes))
}

func checkNilError(err error) { // One way to check for errors instead of writing the if block everytime.
	if err != nil {
		panic(err)
	}
}
