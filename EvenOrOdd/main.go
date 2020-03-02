package main

import "fmt"

func main() {
	nums := []int{3, 6, 8, 10, 5, 7, 2, 1, 4, 9}

	for i, num := range nums {
		if num%2 == 0 {
			fmt.Println(i, "the number is even", num)
		} else {
			fmt.Println(i, "the number is odd", num)
		}
	}
}
