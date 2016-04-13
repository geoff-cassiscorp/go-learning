/*
	This program covers more composition stuff!
*/

package main

import ( 
	"fmt"
)

// a structure in go
type Saiyan struct { 
	Name string
	Power int
	Father *Saiyan
}

func main() {
	gohan := &Saiyan{
  		Name: "Gohan",
  		Power: 1000,
  		Father: &Saiyan {
			Name: "Goku", 
			Power: 9001, 
			Father: nil,
		}, 
	}

	fmt.Printf("%s's father is %s\n", gohan.Name, gohan.Father.Name)
}