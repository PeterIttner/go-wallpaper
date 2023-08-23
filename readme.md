# Wallpaper command line utility

[![BuildAndTest](https://github.com/PeterIttner/go-wallpaper/actions/workflows/ci.yml/badge.svg)](https://github.com/PeterIttner/go-wallpaper/actions/workflows/ci.yml)

## Features:

- Set a random wallpaper from a directory
- Set the bing image of the day wallpaper
- Enhance any wallpaper with the watch word of today
- Can be configured via config file

## Example config files:

### Example 1:
*Take wallpapers from a directory*

```json
{
  "watchWords": {
    "x": 2600,
    "y": 260,
    "fontSize": 40,
    "fontPath": "font/arial.ttf",
    "isActive": true
  },
  "bingFeed": {
    "feedUrl": "https://peapix.com/bing/feed?country=de",
    "isActive": false
  },
  "imageDirectory": {
    "isActive": true,
    "path": "C:/my_wallpaper_images_in_jpeg_format"
  }
}
```

### Example 2:
*Take wallpapers from online service - bing image of the day*

```json
{
  "watchWords": {
    "x": 2600,
    "y": 260,
    "fontSize": 40,
    "fontPath": "font/arial.ttf",
    "isActive": true
  },
  "bingFeed": {
    "feedUrl": "https://peapix.com/bing/feed?country=de",
    "isActive": true
  },
  "imageDirectory": {
    "isActive": false,
    "path": "C:/my_wallpaper_images_in_jpeg_format"
  }
}
```

## For Developers
### Build

In `cmd/wp` directory run:

```bash
go build .
```

### Run

In `cmd/wp` directory run:

```bash
go run wp.exe
```