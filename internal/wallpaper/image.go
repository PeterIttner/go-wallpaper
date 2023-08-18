package wallpaper

type Image struct {
	Title     string `json:"title"`
	Copyright string `json:"copyright"`
	FullURL   string `json:"fullUrl"`
	ThumbURL  string `json:"thumbUrl"`
	ImageURL  string `json:"imageUrl"`
	Date      string `json:"date"`
	PageURL   string `json:"pageUrl"`
}
