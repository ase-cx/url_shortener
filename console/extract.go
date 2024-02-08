package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extract() map[string][]string {
	// Open the file
	file, err := os.Open("template.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Initialize map
	contentMap := make(map[string][]string)

	// Initialize variables for reading
	var currentName string
	var lines []string

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#start#") {
			if currentName != "" {
				contentMap[currentName] = lines
			}
			currentName = strings.TrimPrefix(strings.TrimSuffix(line, "#end#"), "#start#")
			lines = nil
		} else if strings.HasPrefix(line, "#end#") {
			contentMap[currentName] = lines
			currentName = ""
		} else {
			lines = append(lines, line)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Print the map
	// for name, content := range contentMap {
	// fmt.Println(name, ":", content)
	// fmt.Printf("\033[35m")
	// for _, content := range contentMap["logo"] {
	// 	fmt.Printf(content + "\n")
	// }
	return contentMap
}
