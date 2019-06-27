package main

import "fmt"

func main() {
	x := "hello!"
	fmt.Printf("1.- %s\n", x)
	for i := 0; i < len(x); i++ {
		fmt.Printf("2.- %s\n", x)
		x := x[i]
		fmt.Printf("3.- %c\n", x)
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("4.- %c\n", x)
		}
	}
	fmt.Printf("5.- %s\n", x)
}
