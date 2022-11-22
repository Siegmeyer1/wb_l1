package main

import "fmt"

type element struct{}
type set map[string]element

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := set{}
	for _, val := range arr {
		res[val] = element{} // adding to map by existing key overrides value, but don`t add any new elements
	}
	fmt.Println(res)
}
