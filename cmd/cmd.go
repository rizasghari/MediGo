package cmd

import (
	"fmt"
	"log"
	"os"

	ff "github.com/rizasghari/medigo/internal/ffmpeg"
	"github.com/spf13/cobra"
)

const (
	VERSION = "v0.1"
)

var (
	ffmpeg = ff.New()
)

var (
	// Video Thumbnail
	input     string
	output    string
	timestamp int
)

var (
	rootCmd = &cobra.Command{
		Use:   "medigo",
		Short: "MediGo is a ffmpeg based video and audio tool",
		Long:  `Record your screen and audio, create thumbnails, and convert videos to other formats.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Welcome to MediGo!")
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of MediGo",
		Long:  `All software has versions. This is MediGo's!`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("MediGo - Video and Audio Tool v%s\n", VERSION)
		},
	}

	videoThumbnail = &cobra.Command{
		Use:   "thumbnail",
		Short: "Generate video thumbnail",
		Long:  `Generate video thumbnail from video file`,
		Run: func(cmd *cobra.Command, args []string) {
			// log.Println("TODO: Implement video thumbnail command")

			// for arg := range args {
			// 	fmt.Println(args[arg])
			// }
			// log.Printf("input: %s, output: %s, timestamp: %d", input, output, timestamp)

			err := ffmpeg.GenerateThumbnail(input, timestamp, output)
			if err != nil {
				log.Fatal("Error generating thumbnail:", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(videoThumbnail)
}

func setFlags() {
	// Video Thumbnail
	videoThumbnail.Flags().StringVarP(&input, "input", "i", "", "Path to the input file")
	videoThumbnail.MarkFlagRequired("input")

	videoThumbnail.Flags().StringVarP(&output, "output", "o", ".", "Directory to save the output file")
	videoThumbnail.MarkFlagRequired("output")

	videoThumbnail.Flags().IntVarP(&timestamp, "timestamp", "t", 1, "Timestamp (in seconds)")
	videoThumbnail.MarkFlagRequired("timestamp")
}

func Execute() {
	setFlags()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
