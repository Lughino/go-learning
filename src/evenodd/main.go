package main

import "fmt"

func main() {
	sint := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, si := range sint {
		if si%2 == 0 {
			fmt.Println(si, " is even")
		} else {
			fmt.Println(si, " is odd")
		}
	}
}
