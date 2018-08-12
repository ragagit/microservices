package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//Explicitly declaring a "string" variable
	var explicitString string = "Hello World!"

	//Implicitly declaring a "string"
	implicitString := "Hello World!"

	println(explicitString)
	println(implicitString)

	fmt.Println("Variable 'implicitString' is of type:", reflect.TypeOf(implicitString))
	fmt.Println("Variable 'explicitString' is of type:", reflect.TypeOf(explicitString))

	res, err := sum(5, 7)

	println(res, err)

	mul := func(x int, y int) int {
		return x * y
	}
	println(mul(7, 8))

	er := errors.New("Something bad happened")

	if er != nil {
		//panic(er)
		//println(er)
	}

	println(summ(1, 2, 3, 4))
	println(summm(4, 5, 6, 7, 8))

	//Array
	var myArr [50]int
	myArr[0] = 9

	sArr := [3]string{"One", "Two", "Three"}
	println(sArr[0])

	//slices
	mySlice := make([]int, 5)
	mySlice = append(mySlice, 2)
	mySlice[0] = 0 + 5
	mySlice[1] = 1 + 5
	mySlice[2] = 2 + 5
	mySlice[3] = 3 + 5
	mySlice[4] = 4 + 5
	mySlice = append(mySlice, mySlice[:3]...)
	prnSlice(mySlice)
	//To delete the first item
	mySlice1 := mySlice[1:3]
	prnSlice(mySlice1)

	//Maps
	myMap := make(map[string]int)
	myMap["One"] = 1
	myMap["Two"] = 2
	myMap["Three"] = 3

	type MyObject struct {
		Number int    `json:"number"`
		Word   string `json:"string"`
	}

	//	var jsonBytes = []byte(`{"number":5, "string":"packt"}`)
	//	var object MyObject
	//	//err := json.Unmarshal(jsonBytes, &object)
	//	//myJsonMap := make(map[string]interface{})
	//	//jsonData := []byte(`{"hello":"world"}`)
	//	//err := json.Unmarshal(jsonData, &myJsonMap)
	//	//println(myJsonMap["Hello"]
	//
	//	//Uppercase definitions are public
	//	//Lowercase are private
	//
	//	var jsonBlob = []byte(`[
	//    {"Name": "Platypus", "Order": "Monotremata"},
	//    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
	//]`)
	//	type Animal struct {
	//		Name  string
	//		Order string
	//	}
	//	var animals []Animal
	//	//err := json.Unmarshal(jsonBlob, &animals)
	//	if err != nil {
	//		fmt.Println("error:", err)
	//	}
	//	fmt.Printf("%+v", animals)
}

func prnSlice(arr []int) {
	for idx := range arr {
		println(arr[idx])
	}
	println()
}

func summ(args ...int) int {
	res := 0
	for v := range args {
		res = res + v
	}
	return res
}

func summm(args ...int) (result int) {
	for _, v := range args {
		result += v
	}
	return result
}

func sum(x int, y int) (int, int) {
	if x != 0 && y != 0 {
		return (x + y), 0
	} else {
		return 0, -1
	}
}
