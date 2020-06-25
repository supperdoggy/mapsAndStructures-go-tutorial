package main

import "fmt"

func main() {
	// Maps is basically dictionaries
	/*
		to declare: map[keyType]valueType{
			key: value,
		}
	*/

	sortsOfBeer := map[string][]string{
		"Light": {"Bud", "Amstel"},
		"Dark":  {"Lvivske"},
	}
	fmt.Println(sortsOfBeer) // map[Dark:[Lvivske] Light:[Bud Amstel]]

	// To delete key in the map use delete func
	delete(sortsOfBeer, "Light")
	fmt.Println(sortsOfBeer) // map[Dark:[Lvivske]]
	// and if we try to see again what is in Light
	fmt.Println(sortsOfBeer["Light"]) // []
	// it will show empty slice() or 0, but wont add it into the map!
	// to check if variable is valid and not by default there use
	_, ok := sortsOfBeer["Light"]
	fmt.Println(ok)
	// MAPS ARE LINKED

	// Structs, basically this is constructor
	/* declaration
	type *structName* struct{
		name type
		name type
		name type
	}
	*/
	type Doctor struct {
		id         int
		name       string
		companions []string
	}
	var doc1 Doctor = Doctor{id: 3, name: "John", companions: []string{"Dave"}}
	fmt.Println(doc1)
}
