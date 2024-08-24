package ffmpeg

import (
	"fmt"
	"log"
	"os"
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

func (fw FFMPEGWrapper) ListInputDevices() error {
	cmd := exec.Command("ffmpeg", "-f", "avfoundation", "-list_devices", "true", "-i", "\"\"")

	cmd.Env = os.Environ() 
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (fw FFMPEGWrapper) ScreenRecord(video, audio bool) error {
	if utils.IsWindoes() {
		// TODO
	} else if utils.IsLinux() {
		// TODO
	} else if utils.IsMacOS() {
		if !video && !audio {
			return fmt.Errorf("at least one of video or audio must be true")
		}
		// https://ffmpeg.org/ffmpeg-devices.html#avfoundation
		// TODO: Let user choose input devices for video and audio
		videoInput := "0"
		audioInput := "1"
		if !video {
			videoInput = ""
		}
		if !audio {
			audioInput = ""
		}

		cmd := exec.Command("ffmpeg", "-v", "verbose", "-f", "avfoundation", "-framerate", "30", "-i", fmt.Sprintf("%s:%s", videoInput, audioInput), "-pix_fmt", "yuv420p", "./record.mp4")
		log.Printf("cmd: %v", cmd)

		cmd.Env = os.Environ() 
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unsupported os")
	}

	return nil
}
