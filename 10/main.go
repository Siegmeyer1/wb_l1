package main

import "fmt"

// I believe there`s a flaw in result example: for negative values group key is greater than values themselves, while it`s opposite for positives.
// This leads to -9.9 and 9.9 falling into the same group with key 0. Correct example would look like this:
// -30:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -9.9, 9.9}
	m := make(map[int][]float32)
	for _, val := range arr {
		key := int(val) - int(val)%10 // This returns val rounded to multiple of 10
		if val < 0 {                  // Fixing issue with -9.9 and 9.9
			key -= 10
		}
		m[key] = append(m[key], val)
	}
	fmt.Println(m)
}
