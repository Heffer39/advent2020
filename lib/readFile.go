package lib

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(path string) (data []string) {
	abs, err := filepath.Abs(path)
	file, err := os.Open(abs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if len(data) < 1 {
		log.Fatal("data not parsed")
	}
	return
}
