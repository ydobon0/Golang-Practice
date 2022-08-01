package main

import (
	"mymodule/Training"
)

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func main() {
	// Training.Test()
	// Training.BankApp() //Module 6
	//Training.BlackJack() //Module 6
	// Training.GoblinTower() //Module 6
	Training.ErrorHandlingExercise() //Module 7
	//Training.Practice() // concurrencty exercieses
	// Training.July15()
	// Training.July16()

	// p1 := doubleZero{
	// 	person: person{
	// 		First: "James",
	// 		Last:  "Bond",
	// 		Age:   20,
	// 	},

	// 	First:         "Double Zero Seven",
	// 	LicenseToKill: true,
	// }

	// p2 := doubleZero{
	// 	person: person{
	// 		First: "Miss",
	// 		Last:  "MoneyPenny",
	// 		Age:   19,
	// 	},
	// 	First:         "If looks could kill",
	// 	LicenseToKill: false,
	// }

	// // fields and methods of the inner-type are promoted to the outer-type
	// fmt.Println(p1.First, p1.person.First)
	// fmt.Println(p2.First, p2.person.First)
}
