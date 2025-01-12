package main

import (
	"net/http"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func ifElse(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

func loadTextureFromImage(efsImage []byte) rl.Texture2D {
	rlImage := rl.LoadImageFromMemory(getFileExtension(efsImage), efsImage, int32(len(efsImage)))
	return rl.LoadTextureFromImage(rlImage)
}

// GetFileExtension detects the file extension from a byte slice using its MIME type.
func getFileExtension(data []byte) string {
	// Detect the MIME type from the first 512 bytes
	contentType := http.DetectContentType(data)

	// Map MIME type to file extensions
	switch {
	case strings.Contains(contentType, "image/png"):
		return ".png"
	case strings.Contains(contentType, "image/jpeg"):
		return ".jpg"
	case strings.Contains(contentType, "image/gif"):
		return ".gif"
	case strings.Contains(contentType, "image/webp"):
		return ".webp"
	case strings.Contains(contentType, "image/bmp"):
		return ".bmp"
	case strings.Contains(contentType, "text/plain"):
		return ".txt"
	case strings.Contains(contentType, "application/pdf"):
		return ".pdf"
	default:
		return "" // Unknown extension
	}
}
