package main

import _ "embed"

// Embed images
var (
	//go:embed assets/game-controller.png
	iconImage []byte

	//go:embed assets/background.jpg
	backgroundImage []byte

	//go:embed assets/golang.png
	playerImage []byte

	//go:embed assets/NET.png
	netLogoImage []byte

	//go:embed assets/nodejs.png
	nodejsLogoImage []byte

	//go:embed assets/java.png
	javaLogoImage []byte
)
