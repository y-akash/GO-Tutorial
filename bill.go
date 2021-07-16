package main

// here we are declaring our struct
// type structName struct{}			//Syntax
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// // here we are creating new object of bill struct
func newBill(name string) bill {
	// we can also create without provide any value
	// b := bill{}

	// we can also create with provide any one value
	b := bill{
		name: name,
	}

	// we can also create with providing value to all
	// b := bill{
	// 	name:  name,
	// 	items: map[string]float64{},
	// 	tip:   0,
	// }

	return b
}
