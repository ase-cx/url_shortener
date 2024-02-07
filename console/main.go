package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set up a signal channel to catch interrupt signals
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	// spin()
	go donut()

	<-sigc

	// When interrupted, clear the terminal and move cursor to bottom
	moveCursorToBottom()
}

// Function to move the cursor to the bottom of the terminal
func moveCursorToBottom() {
	fmt.Print("\033[9999;1H")
	fmt.Print("\033[?25h") // Show the cursor
	os.Exit(1)
}
