package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// RECIEVER FUNCTION
// here this function is associated with bill struct
// this function can only be called with bill struct object
// we cant directly call this function
// such type of function is called reciever function
// func (variableName structName) functionName() returnType(optional){}		// syntax
// that variableName in this function will be the struct object		*
// we can access all the members of that sturct		*
// but that object will be copy of the orignal struct object		*
// so if we change some member of that sturct than it will not affect to the orignal one		*
func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		// -25 is that variable will take 25 character space
		// for example pie will take 3 character than remaining will be space on the right side
		// if we give 25 that the space will be on the left side
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
	}
	return b
}
