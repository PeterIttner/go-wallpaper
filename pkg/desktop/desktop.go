package desktop

import (
	"github.com/reujab/wallpaper"
)

func ApplyWallpaper(path string) {
	wallpaper.SetFromFile(path)
	wallpaper.SetMode(wallpaper.Crop)
}
