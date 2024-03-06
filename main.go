package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func main() {

	inputPath := flag.String("input", "", "Input JPEG image file path")
	outputPath := flag.String("output", "", "Output JPEG image file path")
	temperatureAdjustment := flag.Int("temperature", 0, "Temperature adjustment value (positive for warmer, negative for cooler)")
	flag.Parse()

	if *inputPath == "" || *outputPath == "" {
		fmt.Println("Error: Input and output file paths are required")
		return
	}

	inputFile, err := os.Open(*inputPath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Decode input image
	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("Error decoding input image:", err)
		return
	}

	adjustedImage := adjustTemperature(inputImage, *temperatureAdjustment)

	outputFile, err := os.Create(*outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, adjustedImage, nil)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	fmt.Println("Image temperature adjustment completed successfully")
}
