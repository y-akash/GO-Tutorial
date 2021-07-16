package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// add item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// update tip
// pointer Reciever
func (b *bill) updateTip(tip float64) {
	// we have dereference the pointer to change the value
	// but go will do that for us in sturct so here we dont have to use astrick *
	// (*b).tip = tip
	b.tip = tip
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}
