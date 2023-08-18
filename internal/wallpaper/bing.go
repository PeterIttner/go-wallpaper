package wallpaper

import (
	"encoding/json"
	"fmt"
	"go-wp/internal/config"
	"io"
	"net/http"
	"os"

	. "github.com/ahmetb/go-linq/v3"
)

func getBingImages(config *config.AppConfig) (error, []Image) {
	res, err := http.Get(config.FeedUrl)
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
	return err, images
}

func filterImagesByDay(images []Image, today string) []Image {
	var imagesOfTody []Image

	From(images).Where(func(i interface{}) bool {
		return i.(Image).Date == today
	}).ToSlice(&imagesOfTody)
	return imagesOfTody
}
