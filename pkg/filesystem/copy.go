package filesystem

import (
	"fmt"
	"os"
)

func Copy(src string, dst string) {

	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Printf("Error reading file %s\n", src)
		os.Exit(1)
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Printf("Error writing file %s\n", dst)
		os.Exit(1)
	}
}
