package main

import (
	"fmt"
)

type somedata struct {
	s string
}

func (d *somedata) printSomething() {
	fmt.Println(d.s)
}

// DoPrintSomething is a exported function
type DoPrintSomething func()

func decorate(f DoPrintSomething) DoPrintSomething {
	return func() { // the return must match the full DoPrintSomething function signature
		fmt.Println("[decorate] before")
		f()
		fmt.Println("[decorate] after")
	}
}

func main() {

	d := somedata{s: "i'm not decorated\n"}
	d.printSomething()

	e := somedata{s: "i now am decorated... woohoo!"}
	decorate(e.printSomething)()

}
