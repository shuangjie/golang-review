package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "/Users/zeng/Documents/gopath/learngo/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
