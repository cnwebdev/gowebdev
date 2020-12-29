// struct exmaple 1
package main

import "fmt"

type person struct {
	fname, lname string
	age          int
}

func main() {
	p1 := person{
		fname: "Chris",
		lname: "Nguyen",
		age:   50,
	}

	p2 := person{
		fname: "Mike",
		lname: "Mc Neilin",
		age:   50,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.fname, p1.lname, p1.age)
	fmt.Println(p2.fname)
}
