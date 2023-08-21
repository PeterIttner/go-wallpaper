package program

import (
	"fmt"
	"go-wp/internal/config"
	"go-wp/internal/drawing"
	"go-wp/internal/wallpaper"
	"go-wp/pkg/desktop"
	"go-wp/pkg/filesystem"
	"go-wp/pkg/watchwords"
	"go-wp/pkg/web"
	"path/filepath"
)

func Run() {
	config := config.Get()
	var imagePath = ""

	if config.BingFeed.IsActive && !config.ImageDirectory.IsActive {
		images := wallpaper.GetBingImages(config.BingFeed.FeedUrl)
		image := wallpaper.FindNewestImage(images)
		imagePath, _ = filepath.Abs("wallpaper.jpg")
		web.DownloadFile(imagePath, image.ImageURL)
	}

	if config.ImageDirectory.IsActive && !config.BingFeed.IsActive {
		images := wallpaper.GetImagesOfDirectory(config.ImageDirectory.Path)
		image := wallpaper.GetRandomImage(images)
		fmt.Printf("Selected wallpaper from directory: %s\n", image)
		filesystem.Copy(image, "wallpaper.jpg")
		imagePath, _ = filepath.Abs("wallpaper.jpg")
	}

	if config.WatchWords.IsActive && imagePath != "" {
		watchwords.UpdateWatchWords()
		watchWords := watchwords.ReadWatchWords()
		watchWord := watchwords.GetWatchWordOfToday(watchWords)
		drawing.DrawTextIntoImage(imagePath, config.WatchWords.FontPath, config.WatchWords.FontSize, config.WatchWords.X, config.WatchWords.Y, watchWord.Lehrtext)
	}

	if imagePath != "" {
		fmt.Printf("Applying wallpaper: %s\n", imagePath)
		desktop.ApplyWallpaper(imagePath)
	}
}
