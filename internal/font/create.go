package font

import (
	_ "embed"
	"fmt"
	"go-wp/pkg/filesystem"
	"os"
	"path"
)

var (
	//go:embed Sketch.ttf
	fontFile []byte
)

func EnsureFont() {

	filename := "fonts/Sketch.ttf"
	filepath := path.Dir(filename)

	if filesystem.Exists(filename) {
		return
	}

	if !filesystem.ExistsDir(filepath) {
		err := os.Mkdir(filepath, 0644)
		if err != nil {
			fmt.Printf("Could not create directory %s, %s\n", filepath, err)
			os.Exit(1)
		}
	}

	err := os.WriteFile(filename, fontFile, 0644)
	if err != nil {
		fmt.Printf("Could not generate font file %s, %s\n", filename, err)
		os.Exit(1)
	}
}
