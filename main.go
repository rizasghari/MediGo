package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
    videoPath := flag.String("video", "", "Path to the input video file")
    timestamp := flag.Int("timestamp", 0, "Timestamp (in seconds) to capture the thumbnail")
    outputDir := flag.String("output", ".", "Directory to save the thumbnail image")

    flag.Parse()

    if *videoPath == "" || *timestamp < 0 {
        log.Fatal("You must specify a valid video path and a non-negative timestamp.")
    }

    outputPath := filepath.Join(*outputDir, "thumbnail.jpg")

    err := GenerateThumbnail(*videoPath, *timestamp, outputPath)
    if err != nil {
        log.Fatal("Error generating thumbnail:", err)
    }

    fmt.Println("Thumbnail generated successfully at", outputPath)
}

func GenerateThumbnail(videoPath string, timestamp int, outputPath string) error {
    cmd := exec.Command("ffmpeg", "-ss", fmt.Sprintf("%d", timestamp), "-i", videoPath, "-vframes", "1", "-q:v", "2", outputPath)

    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("ffmpeg error: %v", err)
    }

    return nil
}