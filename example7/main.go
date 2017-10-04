package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type somedata struct {
	s string
}

func (d *somedata) doesAppThings() string {
	time.Sleep(5 * time.Second)
	return fmt.Sprintf(">>> %v", d.s)
}

func loggingDecorator(f func() string, l *log.Logger) func() string {
	return func() string {
		l.Println("start...")
		defer func() { l.Println("end...") }()
		return f()
	}
}

func main() {

	myLogger := log.New(os.Stdout, "### ", 11)
	s := somedata{s: "yooolooo"}

	sNormal := s.doesAppThings()
	fmt.Println("not decorated", sNormal)

	sDecorated := loggingDecorator(s.doesAppThings, myLogger)()
	fmt.Println("decorated    ", sDecorated)

}
