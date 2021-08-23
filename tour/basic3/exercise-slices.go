package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, uint(dx))
	for i, _ := range image {
		image[i] = make([]uint8, uint(dy))
		for j := 0; j < dy; j++ {
			image[i][j] = uint8((i + j) / 2)
		}

	}
	return image
}

func main() {
	pic.Show(Pic)
}
