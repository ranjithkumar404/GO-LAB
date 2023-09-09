//a go program to find Special Pythagorean triplet
package main

import (
	"fmt"
	
)

func main() {
	var n int
	fmt.Println("Enter the sum")
	fmt.Scanln(&n)

	for a := 1; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				if a+b+c == n && a*a+b*b == c*c {
					fmt.Println("The triplet is", a, b, c)
				}
			}
		}
	}
}
