package main

import ( 
	"fmt"
)

func main() {
	var power int
	power = 9000
	
	fmt.Printf("It's over %d\n", power)

	// is this a comment?

	// This is shorthand for declaration and assignment
	// := can only be used once (declaration)
	//   unless it's a multiple assignment and at least
	//   one member of the list is new (see below)
	// NOTE: Type is inferred from the value on the RHS
	new_power := 9001

	fmt.Printf("It's over %d\n", new_power)

	fmt.Printf("Calling from function now...\n\n")

	newest_power := getPower()

	fmt.Printf("The latest power level is %d\n", newest_power)

	// two things going on here
	// double assignment, but also power is already
	// defined, but since name is new, it's ok by the
	// compiler... yuck.
	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
	// can not change the type of power after declaration though.

	log("this is the log message")

	fmt.Printf("a plus b is: %d\n", add(5,2))

	// power returns 2 values
	value, exists := power_level("goku")

	if exists == false {
  		// handle this error case
		fmt.Printf("Sheeeeeeeeeet! %d\n\n", value)
	}

	// if we don't care to get one of the return values
	_, exists = power_level("goku")

	if exists == false {
  		// handle this error case
  		fmt.Printf("False from single return capture\n\n")
	}
}

func getPower() int { 
	return 9002
}


func log(message string) {
	fmt.Printf("Log message: %s\n", message)
}

// this shorter format is fine if both args are the same type
// func add(a, b int)
func add(a int, b int) int {
	return a + b
}

func power_level(name string) (int, bool) {
	return 9020, false
}

