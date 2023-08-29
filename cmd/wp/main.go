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

var rootCmd = &cobra.Command{
	Use:   "wp",
	Short: "wp is a commandline tool for adjusting your wallpaper",
	Long: `A commandline tools for applying a wallpaper
                either from a web api.
                or from a local directory
                can also print text (watchwords) onto the wallpapers.`,
	Run: func(cmd *cobra.Command, args []string) {
		program.Run()
	},
}

func Init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wp",
	Long:  `All software has versions. This is wp's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("running wp at version: \"%s\", built at %s\n", version, date)
	},
}

func ExecuteRoot() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Init()
	ExecuteRoot()
}
