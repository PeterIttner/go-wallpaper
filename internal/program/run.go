package program

import (
	"fmt"
	"go-wp/internal/config"
	"go-wp/internal/drawing"
	"go-wp/internal/font"
	"go-wp/internal/wallpaper"
	"go-wp/pkg/desktop"
	"go-wp/pkg/filesystem"
	"go-wp/pkg/image"
	"go-wp/pkg/watchwords"
	"go-wp/pkg/web"
	"os"
	"path/filepath"
)

func RunFromConfig() {
	if !config.EnsureConfig() {
		fmt.Printf("Config file created. Please adjust settings.")
		os.Exit(0)
	}
	cfg := config.Get()

	if cfg.BingFeed.IsActive && cfg.ImageDirectory.IsActive {
		fmt.Printf("Only one of the options BingFeed or ImageDirectory can be active at the same time.")
		os.Exit(1)
	}
	Run(cfg.BingFeed.IsActive, cfg.ImageDirectory.IsActive, cfg.WatchWords.IsActive)
}

func Run(useFeed bool, useDirectory bool, useWatchWords bool) {

	cfg := config.Get()
	font.EnsureFont()

	var imagePath = ""

	if useFeed {
		images := wallpaper.GetBingImages(cfg.BingFeed.FeedUrl)
		image := wallpaper.FindNewestImage(images)
		imagePath, _ = filepath.Abs("data/wallpaper.png")
		web.DownloadFile(imagePath, image.ImageURL)
	}

	if useDirectory {
		images := wallpaper.GetImagesOfDirectory(cfg.ImageDirectory.Path)
		image := wallpaper.GetRandomImage(images)
		fmt.Printf("Selected wallpaper from directory: %s\n", image)
		filesystem.Copy(image, "data/wallpaper.jpg")
		imagePath, _ = filepath.Abs("data/wallpaper.jpg")
	}

	if imagePath != "" {
		fmt.Printf("Scale image\n")
		image.Scale(imagePath, "data/wallpaper_resized.png", cfg.Desktop.MaxWidth)

		imagePath, _ = filepath.Abs("data/wallpaper_resized.png")
	}

	if useWatchWords && imagePath != "" {
		watchwords.UpdateWatchWords()
		watchWords := watchwords.ReadWatchWords()
		watchWord := watchwords.GetWatchWordOfToday(watchWords)
		drawing.DrawTextIntoImage(imagePath, cfg.WatchWords.FontPath, cfg.WatchWords.FontSize, cfg.WatchWords.X, cfg.WatchWords.Y, watchWord.Lehrtext)
	}

	if imagePath != "" {
		fmt.Printf("Applying wallpaper: %s\n", imagePath)
		desktop.ApplyWallpaper(imagePath)
	}
}
