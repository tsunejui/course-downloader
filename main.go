package main

import (
	"course-downloader/config"
	"course-downloader/pkg/hiskio"
	"flag"
	"fmt"
	"log"
	"os"
)

var path string

func init() {
	flag.StringVar(&path, "path", "out", "type the path where to store the videos")
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}

func run() error {
	account := os.Getenv("ACCOUNT")
	password := os.Getenv("PASSWORD")

	h, err := hiskio.New(config.Auth{
		Account:  account,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("failed to new hiskio: %v", err)
	}
	return h.Download(path)
}
