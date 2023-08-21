package watchwords

import (
	"encoding/xml"
	"fmt"
	. "github.com/ahmetb/go-linq/v3"
	"io"
	"os"
	"path/filepath"
	"time"
)

func ReadWatchWords() FreeXML {
	xmlFilePath, _ := filepath.Abs(fmt.Sprintf("watchwords/Losungen Free %04d.xml", time.Now().Year()))
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	fileBytes, _ := io.ReadAll(xmlFile)

	var watchWords FreeXML
	xml.Unmarshal(fileBytes, &watchWords)

	return watchWords
}

func GetWatchWordOfToday(data FreeXML) Losungen {
	now := time.Now()

	return From(data.Losungen).Where(func(d interface{}) bool {
		date := fmt.Sprintf("%04d-%02d-%02dT00:00:00", now.Year(), int(now.Month()), now.Day())
		return d.(Losungen).Datum == date
	}).First().(Losungen)

}
