/* Task 1 :-
   create the Person class using a struct
   in Go to represent individuals with attributes
   like name, age, and methods to introduce themselves,
   update their age, and check if they are eligible to vote. */


package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Introduction() {
	fmt.Printf("My name is %s and My age is %d.\n", p.name, p.age)
}

func (p *Person) UpdateAge(newAge int) {
	if newAge > p.age {
		p.age = newAge
	}

}

func (p Person) isEligibleToVote() {
	if p.age >= 18 {
		fmt.Printf("%s is eligible to vote.\n", p.name)
	} else {
		fmt.Printf("%s is not eligible to vote. \n", p.name)
	}
}

func main() {
	Person := Person{name: "sharanu", age: 22}
	Person.Introduction()
	Person.isEligibleToVote()

	Person.UpdateAge(26)
	Person.Introduction()

}


/*
-------Personal Profile System--------
Introduction :-
My name is sharanu and My age is 22.   
---------------------------------------
Who is Eligible to vote :-
sharanu is eligible to vote.
---------------------------------------
Persent Age of Person :-
22
Updated Age of Person
26

*/

