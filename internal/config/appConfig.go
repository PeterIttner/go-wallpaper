package config

type AppConfig struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	FontSize float64 `json:"fontSize"`
	FeedUrl  string  `json:"feedUrl"`
	FontPath string  `json:"fontPath"`
}
