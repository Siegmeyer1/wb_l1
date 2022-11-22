package main

import "fmt"

type element struct{}
type set map[string]element

func main() {
	set1 := set{
		"katana": element{},
		"bow":    element{},
		"honor":  element{},
	}

	set2 := set{
		"katana": element{},
		"honor":  element{},
		"tanto":  element{},
	}

	resultSet := set{}
	for member := range set2 { //iterate over elements of set 2
		if _, ok := set1[member]; ok {
			resultSet[member] = element{} //if this element exists in set 1 too then add it to result set
		}
	}
	fmt.Println(resultSet)
}
