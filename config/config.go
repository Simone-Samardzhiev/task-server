package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadEnvironmentFiles will load all the environment files from a .env file.
func LoadEnvironmentFiles(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening config file: %s", err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		values := strings.Split(line, "=")

		if len(values) != 2 {
			log.Printf("Error parsing config line: %s", line)
		}

		key := strings.TrimSpace(values[0])
		value := strings.TrimSpace(values[1])

		err = os.Setenv(key, value)
		if err != nil {
			log.Printf("Error setting environment variable: %s", err)
		}
	}

	err = file.Close()
	if err != nil {
		log.Printf("Error closing config file: %s", err)
	}
}
