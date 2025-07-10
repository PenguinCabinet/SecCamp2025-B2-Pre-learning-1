package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	Fn := 0
	Fn_minus1 := 1
	Fn_minus2 := 0
	return func() int {
		result := Fn_minus2
		Fn = Fn_minus1 + Fn_minus2
		Fn_minus2 = Fn_minus1
		Fn_minus1 = Fn
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
