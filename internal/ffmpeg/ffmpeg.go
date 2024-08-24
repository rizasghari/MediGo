package ffmpeg

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/rizasghari/medigo/internal/utils"
)

type FFMPEG struct{}

func New() *FFMPEG {
	return &FFMPEG{}
}

func (f *FFMPEG) Convert(videoPath string, outputPath string, outputFormat string) error {
	return nil
}

func (f *FFMPEG) GenerateThumbnail(input string, timestamp int, output string) error {
	output = filepath.Join(output, fmt.Sprintf("thumbnail-%v.jpg", time.Now().Unix()))
    cmd := exec.Command("ffmpeg", "-ss", fmt.Sprintf("%d", timestamp), "-i", input, "-vframes", "1", "-q:v", "2", output)
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("ffmpeg error: %v", err)
    }
    return nil
}

func (f FFMPEG) ScreenRecord() {
    os := utils.GetOS()
    log.Printf("OS: %s", os)
}