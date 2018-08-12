package main

import (
	"fmt"
	"reflect"
)

func main() {
	println("Hello World!")
	var explicit string = "This an explicit vaiable"
	implicit := "This an implicit vaiable"

	fmt.Println(explicit, reflect.TypeOf(implicit))
	fmt.Println(implicit, reflect.TypeOf(explicit))
}
