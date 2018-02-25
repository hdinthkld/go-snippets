package main

import "fmt"

/*
  A map is a type of variable which has the advantage that it does not have a fixed set of fields (like a struct)
  but has the disadvantage that all keys must be of the same type as well as all values having to be the same type.byte

  The keys do not have to be of the same type as the value however.
*/

func main() {
	var m1 map[string]string // This declares a 'nil' map and cannot be used directly

	m1 = make(map[string]string) // The map had to be 'made' in order to be used

	// The keys don't have to exist in order to assign a value to them.
	m1["testkey1"] = "testvalue1"
	m1["testkey2"] = "testvalue2"

	fmt.Println(m1["testkey"]) // This will print out the element value associated with the specified "key"

	// It is also possible to define the entire starting map in one go.
	var m2 = map[string]string{
		"testkey1": "testelement1",
		"testkey2": "testelement2",
	}

	fmt.Println(m2)

	// It is also possible loop over all the elements in a map just like with array or struct
	for k, e := range m1 {
		fmt.Println(k, e)
	}
}
