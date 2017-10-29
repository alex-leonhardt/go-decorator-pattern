package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	counter int64
)

// a custom struct data type
type somedata struct {
	s string
}

// doesAppThings is the method attached to "somedatea"
func (d *somedata) doesAppThings() string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf(">>> %v", d.s)
}

// loggingDecorator outputs to stdout using the custom logger "l"
func loggingDecorator(f func() string, l *log.Logger) func() string {
	return func() string {
		l.Println("start...")
		defer func() { l.Println("end...") }()
		return f()
	}
}

// metricDecorator increments a global "counter" variable
func metricDecorator(f func() string, counter *int64) func() string {
	*counter++
	return func() string {
		defer func() { fmt.Println("count", *counter) }()
		return f()
	}
}

func main() {

	myLogger := log.New(os.Stdout, "### ", 3)
	s := somedata{s: "yooolooo"}

	// count = 1
	sDecorated :=
		loggingDecorator(
			metricDecorator(s.doesAppThings, &counter), myLogger)()
	fmt.Println("decorated    ", sDecorated)

	// count = 2
	sDecorated =
		loggingDecorator( // wraps metricDecorator and adds our custom logger
			metricDecorator(s.doesAppThings, &counter), myLogger)() // wraps s.doesAppThings method, does the counting
	fmt.Println("decorated    ", sDecorated)

	// count = 3
	sDecorated =
		loggingDecorator( // wraps metricDecorator and adds our custom logger
			metricDecorator(s.doesAppThings, &counter), myLogger)() // wraps s.doesAppThings method, does the counting
	fmt.Println("decorated    ", sDecorated)

}
