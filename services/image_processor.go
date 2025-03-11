package services

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"net/http"
	"time"
)

// ImageProcessor is responsible for processing images.
type ImageProcessor struct{}

// NewImageProcessor creates a new instance.
func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}

// ProcessImage downloads the image, calculates its perimeter, and simulates processing delay.
func (ip *ImageProcessor) ProcessImage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to decode image: %v", err)
	}

	// Get image dimensions
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Calculate perimeter
	perimeter := 2 * (width + height)

	// Simulate GPU processing delay (0.1 to 0.4 sec)
	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

	return perimeter, nil
}
