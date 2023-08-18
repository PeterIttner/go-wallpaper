package watchwords

import (
	"encoding/xml"
	"fmt"
	. "github.com/ahmetb/go-linq/v3"
	"io"
	"os"
	"strconv"
)

func ReadWatchWords(xmlFilePath string) *FreeXML {

	xmlFile, err := os.Open(xmlFilePath)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	fileBytes, _ := io.ReadAll(xmlFile)

	var watchWords FreeXML
	xml.Unmarshal(fileBytes, &watchWords)

	return &watchWords
}

func GetWatchWordsOfDay(data FreeXML, year int, month int, day int) Losungen {
	return From(data.Losungen).Where(func(d interface{}) bool {

		monthZeroed := fmt.Sprintf("%02d", month)
		dayZeroed := fmt.Sprintf("%02d", day)

		return d.(Losungen).Datum == strconv.Itoa(year)+"-"+monthZeroed+"-"+dayZeroed+"T00:00:00"
	}).First().(Losungen)

}
