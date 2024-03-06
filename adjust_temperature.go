package main

import (
	"image"
	"image/color"
)

// adjustTemperature adjusts the color tones of the input image based on the temperature adjustment value
func adjustTemperature(input image.Image, adjustment int) image.Image {

	bounds := input.Bounds()
	output := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get pixel color at (x, y)
			pixel := input.At(x, y)

			// Extract RGB components
			r, g, b, _ := pixel.RGBA()

			// Apply temperature adjustment
			r = adjustComponent(r, adjustment)
			b = adjustComponent(b, -adjustment)

			// Set adjusted pixel color in output image
			output.Set(x, y, color.RGBA{
				uint8(r >> 8),
				uint8(g >> 8),
				uint8(b >> 8),
				255,
			})
		}
	}

	return output
}

// adjustComponent adjusts the given color component based on the temperature adjustment value
func adjustComponent(component uint32, adjustment int) uint32 {
	// Apply adjustment while preserving the component's range (0 - 65535)
	adjusted := int(component>>8) + adjustment
	if adjusted < 0 {
		adjusted = 0
	} else if adjusted > 255 {
		adjusted = 255
	}
	return uint32(adjusted) << 8
}
