package main

// Utility functions
func nextFPS(options []int, current int) int {
	for i, fps := range options {
		if fps == current && i < len(options)-1 {
			return options[i+1]
		}
	}
	return options[0]
}

func previousFPS(options []int, current int) int {
	for i, fps := range options {
		if fps == current && i > 0 {
			return options[i-1]
		}
	}
	return options[len(options)-1]
}
