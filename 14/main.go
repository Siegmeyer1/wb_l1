package main

import "fmt"

func main() {
	var integer interface{} = 42
	var str interface{} = "bushido"
	var boolean interface{} = true
	var chann interface{} = make(chan int)

	if i, ok := integer.(int); ok {
		fmt.Println(i)
	}
	if i, ok := str.(string); ok {
		fmt.Println(i)
	}
	if i, ok := boolean.(bool); ok {
		fmt.Println(i)
	}
	if i, ok := chann.(chan int); ok {
		fmt.Println(i)
	}

	// I want to check if you can missinterpr some values
	fmt.Println("|==========================|")

	if i, ok := integer.(string); ok {
		fmt.Println(i)
	}
	if i, ok := integer.(bool); ok {
		fmt.Println(i)
	}
	if i, ok := str.(int); ok {
		fmt.Println(i)
	}
	if i, ok := boolean.(int); ok {
		fmt.Println(i)
	}
}
