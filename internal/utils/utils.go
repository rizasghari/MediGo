package utils

import "runtime"

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
