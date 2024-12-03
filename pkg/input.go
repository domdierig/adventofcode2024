package input

import (
	"bufio"
	"os"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	var lines []string

	// Datei öffnen
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Scanner verwenden, um Zeilen zu lesen
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Fehler während des Scannens überprüfen
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
