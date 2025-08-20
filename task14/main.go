package main

import "reflect"

func printType(obj interface{}) {
	switch obj.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case bool:
		println("bool")
	case chan int:
		println("chan int")
	case chan string:
		println("chan string")
	case chan bool:
		println("chan bool")
	default:
		if reflect.TypeOf(obj).Kind() == reflect.Chan {
			println("chan", reflect.TypeOf(obj).Elem().Name())
			break
		}
		println(reflect.TypeOf(obj).Name())
	}
}

type Imposter struct{}

func main() {
	objs := []any{1, "zxc", true, make(chan int), make(chan Imposter), Imposter{}}
	for _, obj := range objs {
		printType(obj)
	}
}
