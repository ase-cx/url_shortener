package main

import (
	"fmt"
	"time"
)

func animate(contentMap map[string][]string) {
	logo := contentMap["logo"]
	width := len(logo[0])
	line_count := 0
	for {

		for i := 0; i < width; i++ {
			// clearScreen()
			if line_count > 0 {
				fmt.Printf("\033[%dA", line_count+1)
			}
			setRGB(GradientConverter(i * 1))
			for _, line := range logo {
				fmt.Printf("\r\033[K%s\n", line)

			}

			// fmt.Print("\033[0m")
			// fmt.Print("\033[?25l") // Hide the cursor
			time.Sleep(100 * time.Millisecond)
			line_count = len(logo)
			fmt.Printf("%s\033[0m Powered by \033[1m\033[3m\033[4m%s\n\033[0m", spaces(width-25), "ASEnterprise")
		}

	}
}
