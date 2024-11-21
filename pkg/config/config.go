package config

import (
	"bufio"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func LoadEnv(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		eMsg := "error reading .env file"
		err = errors.Wrap(err, eMsg)
		return
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()

		// Skip empty lines and comments
		if strings.TrimSpace(line) == "" || strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}

		// Split by the first '='
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			eMsg := "split by '=' error in .env file"
			err = errors.Wrap(err, eMsg)
			return
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Trim whitespace
		value = strings.Trim(value, `"'`)

		// Set the environment variable
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}
	return nil
}
