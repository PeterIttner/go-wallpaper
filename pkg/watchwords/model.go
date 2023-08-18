package watchwords

//type Welcome7 struct {
//	FreeXML FreeXML `json:"FreeXml"`
//}

type FreeXML struct {
	Losungen                     []Losungen `json:"Losungen"`
	XmlnsXsi                     string     `json:"_xmlns:xsi"`
	XsiNoNamespaceSchemaLocation string     `json:"_xsi:noNamespaceSchemaLocation"`
}

type Losungen struct {
	Datum        string `json:"Datum"`
	Wtag         Wtag   `json:"Wtag"`
	Sonntag      string `json:"Sonntag"`
	Losungstext  string `json:"Losungstext"`
	Losungsvers  string `json:"Losungsvers"`
	Lehrtext     string `json:"Lehrtext"`
	Lehrtextvers string `json:"Lehrtextvers"`
}

type Wtag string

const (
	Dienstag   Wtag = "Dienstag"
	Donnerstag Wtag = "Donnerstag"
	Freitag    Wtag = "Freitag"
	Mittwoch   Wtag = "Mittwoch"
	Montag     Wtag = "Montag"
	Samstag    Wtag = "Samstag"
	Sonntag    Wtag = "Sonntag"
)
