package wallpaper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	. "github.com/ahmetb/go-linq/v3"
)

func GetBingImages(feedUrl string) []Image {
	res, err := http.Get(feedUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	imagesJson, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var images []Image

	json.Unmarshal([]byte(imagesJson), &images)
	if len(images) == 0 {
		fmt.Printf("no bing images available in response: %s\n", res.Body)
		os.Exit(1)
	}

	fmt.Printf("Received %d bing images\n", len(images))
	return images
}

func getImageByDay(images []Image, year int, month int, day int) (Image, error) {
	date := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	image := From(images).FirstWith(func(i interface{}) bool {
		return i.(Image).Date == date
	})

	if image == nil {
		fmt.Printf("No image found for %s\n", date)
		fmt.Printf("Available images: %v\n", images)
		return Image{}, errors.New("")
	}

	return image.(Image), nil
}

func FindNewestImage(images []Image) Image {
	var imageOfToday Image
	var err error
	daySub := 0
	for {
		date := time.Now().AddDate(0, 0, -daySub)
		imageOfToday, err = getImageByDay(images, date.Year(), int(date.Month()), date.Day())
		if err != nil && daySub < 10 {
			daySub++
		}
		if err == nil {
			fmt.Printf("Found an image for today- %ddays\n", daySub)
			break
		}
		if daySub == 10 {
			fmt.Printf("Could not find an image of the last 10 days.\n")
			os.Exit(1)
		}
	}
	return imageOfToday
}
