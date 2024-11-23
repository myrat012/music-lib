package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type DbConfig struct {
	ConnString  string
	MaxConns    int
	ConnTimeout int
}

type AppConfig struct {
	Host string
	Port string
}

type Config struct {
	DBConfig  *DbConfig
	APPConfig *AppConfig
}

func LoadEnv(filename string) (config *Config, err error) {
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
			eMsg := "set env error in .env file"
			err = errors.Wrap(err, eMsg)
			return nil, err
		}
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))

	maxConns, err := strconv.Atoi(os.Getenv("DB_MAXCONNS"))
	if err != nil {
		eMsg := "convert maxconns error in .env file"
		err = errors.Wrap(err, eMsg)
		return nil, err
	}
	conTimeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		eMsg := "convert timeout error in .env file"
		err = errors.Wrap(err, eMsg)
		return nil, err
	}

	dbConf := &DbConfig{
		ConnString:  connStr,
		MaxConns:    maxConns,
		ConnTimeout: conTimeout,
	}
	appConf := &AppConfig{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}

	return &Config{DBConfig: dbConf, APPConfig: appConf}, nil
}
