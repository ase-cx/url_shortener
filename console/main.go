package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	contentMap := extract()

	// // Set up a signal channel to catch interrupt signals
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	// // spin()
	go animate(contentMap)

	<-sigc

	// // When interrupted, clear the terminal and move cursor to bottom
	moveCursorToBottom()
	os.Exit(1)
}

// Function to move the cursor to the bottom of the terminal
func moveCursorToBottom() {
	fmt.Print("\033[9999;1H")
	fmt.Print("\033[0m\033[?25h") // Show the cursor

}
