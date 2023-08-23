package watchwords

import (
	"errors"
	"fmt"
	"go-wp/internal/zip"
	"go-wp/pkg/web"
	"os"
	"time"
)

func UpdateWatchWords() {

	_, err := os.OpenFile(fmt.Sprintf("./watchwords/Losungen Free %04d.xml", time.Now().Year()), os.O_RDONLY, 0644)
	if errors.Is(err, os.ErrNotExist) {
		// handle the case where the file doesn't exist

		url := fmt.Sprintf("https://www.losungen.de/fileadmin/media-losungen/download/Losung_%04d_XML.zip", time.Now().Year())
		err := web.DownloadFile("data/archive.zip", url)
		if err != nil {
			fmt.Printf("Could not download from %s\n", url)
			os.Exit(1)
		}

		err = zip.Unzip("./data/archive.zip", "./watchwords")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Downloaded and extracted watchwords file\n")
	} else {
		fmt.Printf("watchwords file is up to date\n")
	}

}
