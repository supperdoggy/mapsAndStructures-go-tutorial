/*
	Maps
		Collections of value types that are accessed via keys
		Created via literals or via make function
		Members accessed via [key]syntax
			myMap["key"] == "value"
		Check for presence with "value, ok" form of result
		Multiple assognments refer to same underlying data

	Structs
		Collections of disparate data types that describe a single concept
		Keyed by named fields
		Normaly created as types, but anonymous structs are allowed
		Structs are value types
		No inheritance, but can use composition via embedding
		Tags can be added to struct fields to describe field
*/

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
	fmt.Println(doc1.name) // John

	// inheritance in go
	// just write father into kids struct
	/*
		type Person struct{
			name string
		}
		type Doctor struct{
			Person
			Occupation string
		}
	*/
}
