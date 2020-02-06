package main

import (
	"bufio"
	"os"
	"strings"
)

func loadFile(file string) ([]string, error) {
	var lines []string

	loadedFile, err := os.Open(file)
	if err != nil {
		return lines, err
	}
	defer loadedFile.Close()

	scan := bufio.NewScanner(loadedFile)
	for scan.Scan() {
		lines = append(lines, strings.TrimSpace(scan.Text()))
	}
	return lines, nil
}
