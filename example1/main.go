package main

import "fmt"

func decorator(f func(s string)) func(string) {
	return func(s string) {
		fmt.Println("before")
		f(s)
		fmt.Println("after")
	}
}

func printSomething(s string) {
	fmt.Println(s)
}

func main() {

	// without decorator
	printSomething("hello world")

	// with decorator
	decorator(printSomething)("yolo")

}
