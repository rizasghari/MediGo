package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rizasghari/medigo/internal/ffmpeg"
	"github.com/spf13/cobra"
)

const (
	VERSION = "v0.1"
)

var (
	ffmpegWrapper = ffmpeg.New()
)

var (
	// Video Thumbnail
	input     string
	output    string
	timestamp int

	// Record
	video bool
	audio bool
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

	videoThumbnailCmd = &cobra.Command{
		Use:   "thumbnail",
		Short: "Generate video thumbnail",
		Long:  `Generate video thumbnail from video file`,
		Run: func(cmd *cobra.Command, args []string) {
			err := ffmpegWrapper.GenerateThumbnail(input, timestamp, output)
			if err != nil {
				log.Fatal("Error generating thumbnail:", err)
			}
		},
	}

	recordCmd = &cobra.Command{
		Use:   "record",
		Short: "Record your screen and audio",
		Long:  `Record your screen and audio`,
		Run: func(cmd *cobra.Command, args []string) {
			err := ffmpegWrapper.ScreenRecord()
			if err != nil {
				log.Fatal("Error recording:", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(videoThumbnailCmd)
	rootCmd.AddCommand(recordCmd)
}

func setFlags() {
	// Video Thumbnail
	videoThumbnailCmd.Flags().StringVarP(&input, "input", "i", "", "Path to the input file")
	videoThumbnailCmd.MarkFlagRequired("input")
	videoThumbnailCmd.Flags().StringVarP(&output, "output", "o", ".", "Directory to save the output file")
	videoThumbnailCmd.MarkFlagRequired("output")
	videoThumbnailCmd.Flags().IntVarP(&timestamp, "timestamp", "t", 1, "Timestamp (in seconds)")
	videoThumbnailCmd.MarkFlagRequired("timestamp")

	// Record
	recordCmd.Flags().BoolVarP(&video, "video", "v", true, "Record video")
	recordCmd.Flags().BoolVarP(&audio, "audio", "a", true, "Record audio")
	recordCmd.MarkFlagsOneRequired("video", "audio")
}

func Execute() {
	setFlags()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
