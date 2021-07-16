package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func getInput(prompt string, r *bufio.Reader) (string, error) {
// 	fmt.Print(prompt)
// 	input, err := r.ReadString('\n')

// 	return strings.TrimSpace(input), err
// }

func createBill() bill {
	// ************************************************
	// bufio.NewReader(fromWhereYouWantToRead)
	// bufio.NewReader() => this will create a reader
	// fromWhereYouWantToRead => it will tell that from where you want to read
	// here we have mentioned os.Stdin means that we will take it from terminal
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Create a new bill name: ")

	// reader.ReadString('\n') this will return two value
	// first is inputValue and second one will be error
	// in the ReadString method we only pass character not a string
	// tells that read till this character
	// so in this case until you hit enter ie new line till then whichever you write on a terminal will gona read
	name, error := reader.ReadString('\n')
	fmt.Println("Error=>", error)
	name = strings.TrimSpace(name)
	// name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func main() {

	mybill := createBill()
	fmt.Println(mybill)
}
