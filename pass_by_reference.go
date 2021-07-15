package main

import "fmt"

// we are taking pointer thats why we have used as *datatype
func updateName(n *string) {
	*n = "wedge"
}

func main() {
	name := "tifa"

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(m) // using pointer as arg, can pass &name directly instead of "m" as well
	fmt.Println(name)

}

/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|
*/
