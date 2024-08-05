package main

import "fmt"

type myInterface interface {
	getdata() string
}

type myStruct struct {
	data string
}

func (m myStruct) getdata() string {
	return m.data
}

func print(it myInterface) {
	fmt.Println(it.getdata())
}

//func main() {
//	var instance = myStruct{"Hello World"}
//	print(instance)
//}
