package ffmpeg

import (
	"fmt"
	"os/exec"
)

type FFMPEG struct{}

func New() *FFMPEG {
	return &FFMPEG{}
}

func (f *FFMPEG) Convert(videoPath string, outputPath string, outputFormat string) error {
	return nil
}

func (f *FFMPEG) GenerateThumbnail(videoPath string, timestamp int, outputPath string) error {
    cmd := exec.Command("ffmpeg", "-ss", fmt.Sprintf("%d", timestamp), "-i", videoPath, "-vframes", "1", "-q:v", "2", outputPath)

    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("ffmpeg error: %v", err)
    }

    return nil
}