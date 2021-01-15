package configs

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type AppConfigProperties map[string]string

var AbsPath string

func GetAbsPath() {
	dir, err := filepath.Abs("./resources")
	if err != nil {
		log.Fatal(err)
	}

	AbsPath = dir
}

func ReadPropertiesFile(fileName string) (AppConfigProperties, error) {
	config := AppConfigProperties{}

	if len(fileName) == 0 {
		return config, nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}
