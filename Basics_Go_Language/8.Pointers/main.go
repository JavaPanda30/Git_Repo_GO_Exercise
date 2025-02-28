package main

import "fmt"

func main() {
	age := 28
	// this *int is a pointer  type and to create that type add ( & )before a value
	fmt.Print("Age: ")
	agepointer := &age
	fmt.Println(age)
	fmt.Println(agepointer)
	//will give value as pointer to pointer gives original value
	fmt.Println(val(agepointer))
	editAge(agepointer)
	fmt.Println(age)

	
	// pointer func called
	pointer()
}

func val(age *int) int {
	return *age - 18
}

func editAge(age *int) {
	*age = *age + 10
	// no return but will still edit the age directly into the strings
}
