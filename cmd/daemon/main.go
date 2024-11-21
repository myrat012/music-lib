package main

import "github.com/myrat012/test-work-song-lib/pkg/config"

func main() {
	// Load .env
	if err := config.LoadEnv(".env"); err != nil {
		return
	}

}
