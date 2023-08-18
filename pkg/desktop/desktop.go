package desktop

import (
	"github.com/reujab/wallpaper"
)

func ApplyWallpaper(imageOfTodayPath string) {
	wallpaper.SetFromFile(imageOfTodayPath)
	wallpaper.SetMode(wallpaper.Crop)
}
