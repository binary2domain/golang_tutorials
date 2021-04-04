package main

import (
	"fmt"
	"reflect"
)

// var (
// 	name, course string
// 	module       float64
// )

// var (
// 	name   = "Santosh"
// 	course = "Docker Deep Dive"
// 	module = 3.2
// )

func main() {

	// variables declared inside functions and not used raises error
	// however variables declared globally at package level can be unused

	name := "Santosh"
	// course := "Docker"
	module := 3.2
	ptr := &module

	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	// fmt.Println("Course is set to", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Memory address of *module* var is", ptr, "and the value of *module* var is", *ptr)
}
