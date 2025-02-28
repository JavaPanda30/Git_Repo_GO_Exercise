package main

import (
	"fmt"
	"log"
)

// custom data structure named aman
type aman struct {
	name string
}

func main() {
	//custom data type
	var nameVal aman = aman{name: "aman"}
	nameVal.log()

	//sample data input from user
	firstName := GetUserData("Please enter first name: ")
	lastName := GetUserData("Please enter last nmae: ")
	birthdate := GetUserData("Please enter birthdate (MM/DD/YYYY): ")

	//Making Struct
	appUser := Something(firstName, lastName, birthdate)

	Somethingcleaner(appUser)

	SomethingDoneWithPointer(&appUser)

	(appUser).RemoveUserName()

	// direct calling struct

	var appUsercreated2 UserStrct = UserName(firstName, lastName, birthdate)
	Somethingcleaner(appUsercreated2)

	var appUsercreated3 *UserStrct = UserNamePointer(firstName, lastName, birthdate)
	Somethingcleaner(*appUsercreated3)

	//validate func

	var UserChkd *UserStrct

	UserChkd, err := ValidatingUserCreated(firstName, lastName, birthdate)

	if err == nil {
		Somethingcleaner(*UserChkd)
	} else {
		log.Fatalf("User data invalid")
	}
}

func (n aman) log() {
	fmt.Println(n.name)
}
