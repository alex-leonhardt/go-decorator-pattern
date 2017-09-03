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

func decorate(f func()) func() {
	return func() {
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
