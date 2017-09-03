package main

import (
	"fmt"
)

type somedata struct {
	s string
}

func (d *somedata) printSomething(prefix string) {
	fmt.Println(prefix, d.s)
}

// DoPrintSomething is a exported function
type DoPrintSomething func(string)

func decorate(f DoPrintSomething) DoPrintSomething {
	return func(s string) { // the return must match the full DoPrintSomething function signature
		fmt.Println("[decorate] before")
		f(s)
		fmt.Println("[decorate] after")
	}
}

func main() {

	d := somedata{s: "i'm not decorated\n"}
	d.printSomething(">>")

	e := somedata{s: "i now am decorated... woohoo!"}
	decorate(e.printSomething)("==>") // <-- notice the ending () .. this is where function args would be going (we'll be doing that later)

}
