package image

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
)

func Scale(inputImagePath string, outputImagePath string, maxWidth uint) {
	file, err := os.Open(inputImagePath)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(maxWidth, 0, img, resize.Lanczos3)

	out, err := os.Create(outputImagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
