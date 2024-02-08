package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

const (
	width  = 40
	height = 20
)

// GradientConverter converts a value from 0 to 100 to an RGB color gradient
func GradientConverter(value int) (int, int, int) {
	defer func() {
		if r := recover(); r != nil {
			moveCursorToBottom()
			fmt.Println(r)
			fmt.Println(string(debug.Stack()))
			os.Exit(0)

		}
	}()
	// Ensure value is within 0 to 100 range
	if value < 0 {
		value = 0
	} else if value > width {
		value = width
	}

	// Define RGB color points along the rainbow spectrum
	// Red, Orange, Yellow, Green, Blue, Indigo, Violet
	colors := [][]int{
		{255, 0, 0},   // Red
		{255, 165, 0}, // Orange
		{255, 255, 0}, // Yellow
		{0, 255, 0},   // Green
		{0, 0, 255},   // Blue
		{75, 0, 130},  // Indigo
		{128, 0, 128}, // Violet
	}

	// Calculate the number of color segments
	numSegments := len(colors) - 1
	// Calculate the width of each segment
	segmentWidth := width / float64(numSegments)

	// Find the segment that contains the input value
	segment := value / int(segmentWidth)
	// Calculate the position within the segment
	positionInSegment := float64(value%int(segmentWidth)) / segmentWidth

	// Get the start and end colors for the current segment
	startColor := colors[segment]
	endColor := colors[(segment+1)%numSegments]

	// Interpolate between start and end colors
	r := interpolate(startColor[0], endColor[0], positionInSegment)
	g := interpolate(startColor[1], endColor[1], positionInSegment)
	b := interpolate(startColor[2], endColor[2], positionInSegment)

	return r, g, b
}

// interpolate calculates the interpolated value between two numbers
func interpolate(start, end int, weight float64) int {
	return int(float64(end-start)*weight) + start
}

func setColor(color int) {
	// int from 30 to 37
	if color < 30 || color > 37 {
		color %= 8
		color += 30
		fmt.Printf("\033[%dm", color)
	}
}

func setRGB(r, g, b int) {
	fmt.Printf("\033[38;2;%d;%d;%dm", r, g, b)
}

func donut() {
	dots := []string{".", "..", "..."}
	frames := []string{
		"    ######    ",
		"  ##      ##  ",
		"##          ##",
		"##          ##",
		"##          ##",
		"##          ##",
		"  ##      ##  ",
		"    ######    ",
	}
	frame_count := 0
	for {
		for i := 0; i < width; i++ {
			// clearScreen()
			if frame_count > 0 {
				fmt.Printf("\033[%dA", frame_count)
			}
			setRGB(GradientConverter(i * 1))
			for _, line := range frames {
				fmt.Printf("\r\033[K%s%s\n", spaces(i), line)

			}

			dot_with_color := fmt.Sprintf("\033[%dm%s \033[0m", 30+i%18, dots[i/5%3])

			fmt.Print("\033[0m")
			fmt.Printf("\r\033[Kthis is a running donut %s\n", dot_with_color)

			// fmt.Print("\033[?25l") // Hide the cursor
			time.Sleep(100 * time.Millisecond)
			frame_count = len(frames) + 1
		}
	}
}

// func clearScreen() {
// 	// fmt.Print("\033[H\033[2J")
// }

func spaces(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}
