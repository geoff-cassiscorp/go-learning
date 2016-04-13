package main

import ( 
	"fmt"
)

// a structure in go
type Saiyan struct { 
	Name string
	Power int
}

func main() {
	// The following are all valid ways to
	// do assignment with a struct
	goku := Saiyan {
  		Name: "Goku",
  		Power: 9000, // trailing comma on last member is required!
	}

	fmt.Printf("%s has a power level of %d\n", goku.Name, goku.Power)

	// this is cool too
	vegita := Saiyan{}
	vegita.Name = "Vegita"


	trunks := Saiyan{Name: "Trunks"}
	trunks.Power = 4000

	fmt.Printf("%s\n", vegita.Name)
	fmt.Printf("%s\n", trunks.Name)

	// can also skip field names, but could be confusing with many members
	//goku := Saiyan{"Goku", 9000}

	// setting our var to the memory address of a Saiyan obj
	son_goku := &Saiyan{"Son Goku", 10000}

	fmt.Printf("%s's power is %d\n", son_goku.Name, son_goku.Power)
	Super(son_goku)
	fmt.Printf("%s's power is %d\n", son_goku.Name, son_goku.Power)

	// now let's try a method on Saiyan
	mr_huge := &Saiyan{"Mr. Huge", 22000}
	mr_huge.Huge()
	fmt.Printf("%s has a power level of %d\n\n", mr_huge.Name, mr_huge.Power)


	// you can do this either way.. first way is cleaner
	goku2 := new(Saiyan)
	goku2.Name = "goku"
	goku2.Power = 9001
	//vs
	goku3 := &Saiyan {
  		Name: "goku",
  		Power: 9000,
	}

	fmt.Printf("goku2: %s\n", goku2.Name)
	fmt.Printf("goku3: %s\n", goku3.Name)

}

// This function expects a pointer to a Saiyan obj
func Super(s *Saiyan) {
	s.Power += 10000
}

// the type *Saiyan is the RECEIVER of the Huge method
// this acts similarly to class methods in C++
func (s *Saiyan) Huge() {
	s.Power += 80000
}


// structures don't have constructors, instead you create a function
// that returns an instance of the desired type (like a factory)
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name: name,
    	Power: power,
  	}
}

// it doesn't have to return a pointer, this is also totally fine
/*
func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name: name,
    	Power: power,
  	}
}
*/