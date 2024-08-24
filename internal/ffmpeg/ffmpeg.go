package ffmpeg

import (
	"fmt"
	"os/exec"

	"github.com/rizasghari/medigo/internal/utils"
)

type FFMPEGWrapper struct{}

func New() *FFMPEGWrapper {
	return &FFMPEGWrapper{}
}

func (fw *FFMPEGWrapper) Convert(videoPath string, outputPath string, outputFormat string) error {
	return nil
}

func (fw *FFMPEGWrapper) GenerateThumbnail(input string, timestamp int, output string) error {
	cmd := exec.Command("ffmpeg", "-ss", fmt.Sprintf("%d", timestamp), "-i", input, "-vframes", "1", "-q:v", "2", output)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (fw FFMPEGWrapper) ScreenRecord() error {
	if utils.IsWindoes() {
		// TODO
	} else if utils.IsLinux() {
		// TODO
	} else if utils.IsMacOS() {
		// cmd := exec.Command("ffmpeg", "-f", "avfoundation", "-list_devices", "true", "-i", "\"\"")
        cmd := exec.Command("ffmpeg", "-f", "avfoundation", "-i", "2:1", "output.mkv")
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unsupported os")
	}

	return nil
}
