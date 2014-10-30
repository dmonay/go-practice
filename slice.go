package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	fmt.Println(pic)
	for index := range pic {
		row := make([]uint8, dx)
		for index2 := range row {
			row[index2] = uint8((index2 + index) / 2)
		}
		pic[index] = row
	}
	return pic
}

func main() {
	fmt.Println(Pic(2, 3))
}
