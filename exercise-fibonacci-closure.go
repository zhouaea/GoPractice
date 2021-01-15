package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	number := 0
	im2 := 0
	im1 := 0
	
	return func() int {
		switch i {
			case 0: 
				number = 0
			case 1:
				number = 1
				im1 = number
			default: 
				number = im1 + im2
				im2 = im1
				im1 = number		
		}
		i++
		return number
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
