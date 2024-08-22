package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
    // Define command-line flags
    videoPath := flag.String("video", "", "Path to the input video file")
    timestamp := flag.Int("timestamp", 0, "Timestamp (in seconds) to capture the thumbnail")
    outputDir := flag.String("output", ".", "Directory to save the thumbnail image")

    // Parse the flags
    flag.Parse()

    // Validate input
    if *videoPath == "" || *timestamp < 0 {
        log.Fatal("You must specify a valid video path and a non-negative timestamp.")
    }

    // Generate the output path for the thumbnail
    outputPath := filepath.Join(*outputDir, "thumbnail.jpg")

    // Call the thumbnail generation function
    err := GenerateThumbnail(*videoPath, *timestamp, outputPath)
    if err != nil {
        log.Fatal("Error generating thumbnail:", err)
    }

    fmt.Println("Thumbnail generated successfully at", outputPath)
}

// GenerateThumbnail generates a thumbnail image from the given video file at the specified timestamp.
func GenerateThumbnail(videoPath string, timestamp int, outputPath string) error {
    // Build the FFmpeg command
    cmd := exec.Command("ffmpeg", "-ss", fmt.Sprintf("%d", timestamp), "-i", videoPath, "-vframes", "1", "-q:v", "2", outputPath)

    // Execute the command and check for errors
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("ffmpeg error: %v", err)
    }

    return nil
}