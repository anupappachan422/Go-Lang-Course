package main

import "fmt"

func main() {

	slc := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range slc {
		if v%2 == 0 {
			fmt.Println(v, "even")
		} else {
			fmt.Println(v, "odd")
		}
	}

}
