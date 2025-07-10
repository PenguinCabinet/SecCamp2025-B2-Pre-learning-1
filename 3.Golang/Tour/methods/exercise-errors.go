package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	result, err := Sqrt(2)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	result, err = Sqrt(-2)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
