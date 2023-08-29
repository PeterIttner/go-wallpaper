# Wallpaper command line utility

[![BuildAndTest](https://github.com/PeterIttner/go-wallpaper/actions/workflows/ci.yml/badge.svg)](https://github.com/PeterIttner/go-wallpaper/actions/workflows/ci.yml)

<img src="logo.jpg" width="350">

## Features:

- Set a random wallpaper from a directory
- Set the bing image of the day wallpaper (4K resolution)
- Enhance any wallpaper with the watch word of today
- Downloads current watchwords of the year if not present
- Scales image to fixed width and keeps aspect ratio
- Position, fontsize, urls, etc. can be configured via config file
- Config file will be generated, if not present 

## Usage

1. Download latest version from [releases](https://github.com/PeterIttner/go-wallpaper/releases).
2. Unpack archieve and execute go-wallpaper[.exe] executable
3. Adjust the default configuration file
4. (Optional) Create a shortcut for your OS startup mechanism


## Example command line parameters
To override the config file behaviour you can execute the executable with the following commandline arguments:

```shell
# Print help
go-wallpaper --help

# Print the current version of the executable
go-wallpaper version

# Force using the directory from the config file as image source
go-wallpaper dir

# Force using the directory from the config file as image source
go-wallpaper bing

# Add this parameter with the previous commands to optionally enable printing the watchwords on top of the wallpaper 
--watchwords

# e.g.
go-wallpaper dir --watchwords

```

## Example config files:

### Example 1:
*Take wallpapers from a directory*

```json
{
  "watchWords": {
    "x": 1500,
    "y": 1000,
    "fontSize": 50,
    "fontPath": "fonts/Sketch.ttf",
    "isActive": true
  },
  "bingFeed": {
    "feedUrl": "https://peapix.com/bing/feed?country=de",
    "isActive": false
  },
  "imageDirectory": {
    "isActive": true,
    "path": "C:/tools/wp/images"
  },
  "desktop": {
    "maxWidth": 2560
  }
}
```

### Example 2:
*Take wallpapers from online service - bing image of the day*

```json
{
  "watchWords": {
    "x": 1500,
    "y": 1000,
    "fontSize": 50,
    "fontPath": "fonts/Sketch.ttf",
    "isActive": true
  },
  "bingFeed": {
    "feedUrl": "https://peapix.com/bing/feed?country=de",
    "isActive": true
  },
  "imageDirectory": {
    "isActive": false,
    "path": "C:/tools/wp/images"
  },
  "desktop": {
    "maxWidth": 2560
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
go run .
```

### Github Release with GoReleaser

[GoReleaser](https://goreleaser.com/)  is used to generate releases for different OS and architectures in github releases.

[GoReleaser](https://goreleaser.com/) is triggered on every git-tag that starts with `v*`, e.g. `v1.0.0`.

### Useful information

Project inspiration: [https://github.com/PeterIttner/BingWallpaper](https://github.com/PeterIttner/BingWallpaper)

#### Links

- [https://www.losungen.de/digital/daten](https://www.losungen.de/digital/daten)
- [https://bing.gifposter.com/](https://bing.gifposter.com/)
- [https://peapix.com/bing/feed?country=de](https://peapix.com/bing/feed?country=de) 4K Image API
