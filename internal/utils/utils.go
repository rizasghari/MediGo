package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func GetOS() string {
	switch runtime.GOOS {
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default:
		return runtime.GOOS
	}
}

func IsWindoes() bool {
	return GetOS() == "windows"
}

func IsLinux() bool {
	return GetOS() == "Linux"
}

func IsMacOS() bool {
	return GetOS() == "macOS"
}

func GetCodecOptions(inputFile string) (vcodec, acodec, formatOpt, outputFile string) {
    // Get the file extension without the dot
    ext := strings.ToLower(filepath.Ext(inputFile))
    ext = strings.TrimPrefix(ext, ".")

    // Determine codecs and format options based on the file extension
    switch ext {
    case "mp4":
        vcodec = "libx264"
        acodec = "aac"
        formatOpt = ""
    case "mov":
        vcodec = "libx264"
        acodec = "aac"
        formatOpt = "-f mov"
    case "webm":
        vcodec = "libvpx-vp9"
        acodec = "libopus"
        formatOpt = ""
    case "mkv":
        vcodec = "libx265"
        acodec = "libopus"
        formatOpt = ""
    case "avi", "flv":
        vcodec = "libx264"
        acodec = "aac"
        formatOpt = "-f mp4"
        outputFile = GetFilePathWithoutExtension(inputFile) + "_compressed.mp4"
    default:
        fmt.Println("WARNING: unsupported video format - trying with default codecs")
        vcodec = "libx264"
        acodec = "aac"
        formatOpt = ""
        outputFile = GetFilePathWithoutExtension(inputFile) + "_default.mp4"
    }

    // Default output file name if not set
    if outputFile == "" {
        outputFile = GetFilePathWithoutExtension(inputFile) + "_compressed" + filepath.Ext(inputFile)
    }

    return vcodec, acodec, formatOpt, outputFile
}

func GetFilePathWithoutExtension(filePath string) string {
    return strings.TrimSuffix(filePath, filepath.Ext(filePath))
}