package main

import (
	"fmt"
	"time"
)

type somedata struct {
	s string
}

func (d *somedata) returnStringWithPrefix(prefix string) string {
	return fmt.Sprintf("%s %s", prefix, d.s) // Sprintf returns a formatted string, rather than trying to try output it somewhere
}

// DoPrintSomething is a exported function
type DoPrintSomething func(string) string

func decorate(f DoPrintSomething) DoPrintSomething {
	return func(s string) string { // the return must match the full DoPrintSomething function signature
		start := time.Now()
		defer func() { fmt.Println("[decorate] took:", time.Since(start)) }() // as we return f(s) we must use defer as this executes "after" a function has "returned"
		return f(s)
	}
}

func main() {

	d := somedata{s: "i'm not decorated\n"}
	ds := d.returnStringWithPrefix(">>")
	fmt.Println(ds)

	e := somedata{s: "i now am decorated... woohoo!"}
	es := decorate(e.returnStringWithPrefix)("==>")
	fmt.Println(es)

}
