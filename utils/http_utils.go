package utils

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func ProcessImage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read image data: %v", err)
	}

	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
	return nil
}
