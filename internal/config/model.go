package config

type AppConfig struct {
	WatchWords struct {
		X        float64 `json:"x"`
		Y        float64 `json:"y"`
		FontSize float64 `json:"fontSize"`
		FontPath string  `json:"fontPath"`
		IsActive bool    `json:"isActive"`
	} `json:"watchWords"`
	BingFeed struct {
		FeedUrl  string `json:"feedUrl"`
		IsActive bool   `json:"isActive"`
	} `json:"bingFeed"`
	ImageDirectory struct {
		IsActive bool   `json:"isActive"`
		Path     string `json:"path"`
	} `json:"imageDirectory"`
	Desktop struct {
		MaxWidth uint `json:"maxWidth"`
	} `json:"desktop"`
}
