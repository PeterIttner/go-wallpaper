package wallpaper

import (
	"fmt"
	"go-wp/internal/config"
	"go-wp/internal/drawing"
	"go-wp/pkg/desktop"
	"go-wp/pkg/watchwords"
	"go-wp/pkg/web"
	"path/filepath"
	"strconv"
	"time"
)

func ApplyBingWallpaperOfTheDay() {
	config := config.Get()

	now := time.Now()
	today := now.Format("2006-01-02")
	_, images := getBingImages(config)
	imagesOfToday := filterImagesByDay(images, today)

	imageOfTodayPath, _ := filepath.Abs("wallpaper.jpg")
	web.DownloadFile(imageOfTodayPath, imagesOfToday[0].ImageURL)

	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	watchwordsfile, _ := filepath.Abs("watchwords/Losungen Free " + strconv.Itoa(year) + ".xml")
	data := watchwords.ReadWatchWords(watchwordsfile)
	word := watchwords.GetWatchWordsOfDay(*data, year, month, day)

	fmt.Println(word.Lehrtextvers)
	fmt.Println(word.Lehrtext)

	drawing.DrawTextIntoImage(imageOfTodayPath, config.FontPath, config.FontSize, config.X, config.Y, word.Lehrtext)

	desktop.ApplyWallpaper(imageOfTodayPath)
}
