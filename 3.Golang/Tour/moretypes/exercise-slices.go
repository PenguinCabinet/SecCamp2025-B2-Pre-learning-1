package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, 256)
	for y := 0; y < 256; y++ {
		result[y] = make([]uint8, 256)
		for x := 0; x < 256; x++ {
			result[y][x] = uint8((x + y) / 2)
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
