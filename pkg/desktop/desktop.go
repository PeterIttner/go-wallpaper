package desktop

import (
	"fmt"
	"github.com/reujab/wallpaper"
	"os"
)

func ApplyWallpaper(path string) {
	err := wallpaper.SetFromFile(path)
	if err != nil {
		fmt.Printf("Error applying wallpaper: %s, %s\n", path, err)
		os.Exit(1)
	}

	err = wallpaper.SetMode(wallpaper.Crop)
	if err != nil {
		fmt.Printf("Error setting wallpaper mode: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Wallpaper applied")
}
