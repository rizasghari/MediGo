package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "medigo",
	Short: "MediGo is a ffmpeg based video and audio tool",
	Long: `Record your screen and audio, create thumbnails, and convert videos to other formats.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome to MediGo!")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of MediGo",
	Long:  `All software has versions. This is MediGo's`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("MediGo - Video and Audio Tool v0.1")
	},
  }

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
