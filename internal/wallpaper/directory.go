package wallpaper

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GetImagesOfDirectory(directory string) []string {
	entries, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err)
		os.Exit(1)
	}

	var images []string
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("Error reading entry: %s\n", entry)
			os.Exit(1)
		}

		if !info.IsDir() {

			if strings.HasSuffix(info.Name(), ".jpg") || strings.HasSuffix(info.Name(), ".jpeg") {
				path, _ := filepath.Abs(fmt.Sprintf("%s/%s", directory, info.Name()))
				images = append(images, path)
			}

		}
	}

	if len(images) == 0 {
		fmt.Printf("No images found in directory %s\n", directory)
		os.Exit(1)
	}

	return images
}

func GetRandomImage(images []string) string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return images[r.Intn(len(images))]
}
