// anonymous struct example 1
package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	fmt.Println(p1)
}
