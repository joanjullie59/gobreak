package main

import "fmt"

func main() {
	for i := 1; i <= 8; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}
