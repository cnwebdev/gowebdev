// embeded structs example 1
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	job    string
	salary int
	ltk    bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   42,
		},
		ltk:    true,
		job:    "Agent Bond",
		salary: 1,
	}

	fmt.Println(p1)
	fmt.Println(sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.job)
	fmt.Println(p1.first, p1.last, p1.age)
}
