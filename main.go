package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {

	// Name Matrix
	name := [][]int{
		0: {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		1: {0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		2: {0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		3: {0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1},
		4: {0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0},
		5: {0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1},
		6: {0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		7: {0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1},
		8: {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	// Pixel Config
	pixelWidth := 100
	pixelHeight := pixelWidth
	pixelColor := color.Black

	width := len(name[0]) * pixelWidth
	height := len(name) * pixelHeight

	// Image we are going to draw
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)

	// White background
	draw.Draw(img, rect, image.NewUniform(color.White), image.ZP, draw.Over)

	// Draw Image
	for i, row := range name {
		for j, pixel := range row {
			if pixel == 1 {
				// Coordinates
				x := j * pixelWidth
				y := i * pixelHeight

				// Position
				xpos := x + pixelWidth
				ypos := y + pixelHeight

				// Draw
				draw.Draw(
					img,
					image.Rect(x, y, xpos, ypos),
					image.NewUniform(pixelColor),
					image.ZP,
					draw.Over,
				)
			}
		}
	}

	// Create file to save the image
	f, err := os.Create("luispixel.png")
	if err != nil {
		panic(err)
	}

	// Save image
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

	f.Close()
	fmt.Println("Open luispixel.png")
}
