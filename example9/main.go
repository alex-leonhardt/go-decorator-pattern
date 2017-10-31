package main

import (
	"fmt"
	"log"
	"os"
)

// Echoer is a interface that consists of a single method "Echo". Echo is a method that may return an error or nil.
type Echoer interface {
	Echo() error
	Dump() string
}

// a custom struct data type
type text struct {
	s string
}

// Here we have a method Echo() that implements the 1st part of the Echoer interface, the implementation is implicit
func (t text) Echo() error {
	_, err := fmt.Println(t.s)
	return err
}

// Here we have a method Dump() that implements the 2nd part of the Echoer interface, the implementation is implicit and now complete
func (t text) Dump() string {
	return fmt.Sprintf(t.s)
}

func printIt(e Echoer) {
	e.Echo()
}

func dumpIt(e Echoer) string {
	return e.Dump()
}

// Decorates the Echoer interface.
func loggingDecorator(e Echoer, l *log.Logger) Echoer {
	l.Println(">>>", dumpIt(e))
	return e
}

func main() {

	myLogger := log.New(os.Stdout, "### ", 3)
	t := text{s: "happy halloween"}

	// not decorated, sends t, which has method Echo, which implements Echoer
	printIt(t)

	tD := loggingDecorator(t, myLogger) // with the dumpIt function, or calleing e.Dump() we can access the value of the struct and output it as a log
	tD.Echo()                           // not in the "log" as the execution of Echo() happened after the interface was already executed

}
