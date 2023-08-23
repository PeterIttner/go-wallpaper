package filesystem

import (
	"fmt"
	"os"
	"path"
)

func Copy(src string, dst string) {

	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Printf("Error reading file %s, %s\n", src, err)
		os.Exit(1)
	}

	if !ExistsDir(path.Dir(dst)) {
		err = os.Mkdir(path.Dir(dst), 0644)
		if err != nil {
			fmt.Printf("Error creating directory file %s, %s\n", path.Dir(dst), err)
			os.Exit(1)
		}
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Printf("Error writing file %s, %s\n", dst, err)
		os.Exit(1)
	}
}
