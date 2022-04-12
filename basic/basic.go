package main

import "fmt"

var a = 1

func variableZeroValue() {
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var b int = 4
	//a = 3
	var s string = "abc"
	fmt.Println(a, b, s)
}

func Sqrt() {

}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
}
