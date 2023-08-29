package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-wp/internal/program"
	"os"
)

var (
	version = "dev"
	date    = "unknown"
)

var useWatchWordsFlag bool

func main() {
	Init()
	ExecuteRoot()
}

func Init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(bingCmd)
	rootCmd.AddCommand(dirCmd)

	bingCmd.PersistentFlags().BoolVar(&useWatchWordsFlag, "watchwords", false, "Apply the watchwords watermark the wallpaper")
	dirCmd.PersistentFlags().BoolVar(&useWatchWordsFlag, "watchwords", false, "Apply the watchwords watermark the wallpaper")

	rootCmd.DisableSuggestions = false
	rootCmd.SuggestionsMinimumDistance = 1
}

var rootCmd = &cobra.Command{
	Use:   "wp",
	Short: "wp is a commandline tool for adjusting your wallpaper",
	Long: `A commandline tools for applying a wallpaper
                either from a web api.
                or from a local directory
                can also print text (watchwords) onto the wallpapers.`,
	Run: func(cmd *cobra.Command, args []string) {
		program.RunFromConfig()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wp",
	Long:  `All software has versions. This is wp's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("running wp at version: \"%s\", built at %s\n", version, date)
	},
}

var bingCmd = &cobra.Command{
	Use:   "bing",
	Short: "Use the bing feed to apply the wallpaper",
	Long: `The bing feed will provide the wallpaper of the day
           which will be applied. This overrides the config file settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		program.Run(true, false, useWatchWordsFlag)
	},
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Use the images from a directory to apply the wallpaper",
	Long: `The images from the directory will used to apply the wallpaper.
           This overrides the config file settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		program.Run(false, true, useWatchWordsFlag)
	},
}

func ExecuteRoot() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
